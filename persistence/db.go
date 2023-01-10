package persistence

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/localization"
	"os"
	"path"
	"time"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

const (
	ruleSetPersistenceFileName   = "rule-set-persistence.psc"
	cardPersistenceFileName      = "card-persistence.psc"
	characterPersistenceFileName = "character-persistence.psc"
	skillPersistenceFileName     = "skill-persistence.psc"
	sqlite3DBFileName            = "gisb-sqlite3.db"
)

var (
	storagePath = ""
	loaded      = false
)

var (
	RuleSetPersistence   = newFactoryPersistence[RuleSet]()
	CardPersistence      = newFactoryPersistence[Card]()
	CharacterPersistence = newFactoryPersistence[Character]()
	SkillPersistence     = newFactoryPersistence[Skill]()

	LocalizationPersistence = newMemoryCache[string, localization.LanguagePack]()

	CardDeckPersistence DatabasePersistence[uint, CardDeck]
	PlayerPersistence   DatabasePersistence[uint, Player]
)

// SetStoragePath 设置持久化文件的存放位置
func SetStoragePath(path string) error {
	if s, err := os.Stat(path); err == nil && s.IsDir() {
		storagePath = path
		return nil
	} else {
		return fmt.Errorf("path %s is not a directory or is not exist", path)
	}
}

// Serve 开启持久化模块的服务
func Serve(flushFeq time.Duration, errChan chan error) {
	RuleSetPersistence.Serve(flushFeq, storagePath, ruleSetPersistenceFileName, errChan)
	CardPersistence.Serve(flushFeq, storagePath, cardPersistenceFileName, errChan)
	CharacterPersistence.Serve(flushFeq, storagePath, characterPersistenceFileName, errChan)
	SkillPersistence.Serve(flushFeq, storagePath, skillPersistenceFileName, errChan)
}

// Load 从持久化文件读取信息，写入持久化模块
func Load(errChan chan error) {
	if !loaded {
		var err error

		// 初始化Sqlite3
		if sqlite3DB, err = xorm.NewEngine("sqlite3", path.Join(storagePath, sqlite3DBFileName)); err != nil {
			errChan <- err
		} else {
			sqlite3DB.SetMapper(core.SameMapper{})
			if err = sqlite3DB.Sync2(Player{}, CardDeck{}); err != nil {
				errChan <- err
			}

			var success bool
			if success, CardDeckPersistence = newDatabasePersistence[uint, CardDeck](errChan); !success {
				errChan <- fmt.Errorf("failed to create database factoryPersistence with entity: %+v", CardDeck{})
			}

			if success, PlayerPersistence = newDatabasePersistence[uint, Player](errChan); !success {
				errChan <- fmt.Errorf("failed to create database factoryPersistence with entity: %+v", Player{})
			}
		}

		// 初始化Factories
		{
			if err = RuleSetPersistence.Load(path.Join(storagePath, ruleSetPersistenceFileName)); err != nil {
				errChan <- err
			}

			if err = CardPersistence.Load(path.Join(storagePath, cardPersistenceFileName)); err != nil {
				errChan <- err
			}

			if err = CharacterPersistence.Load(path.Join(storagePath, characterPersistenceFileName)); err != nil {
				errChan <- err
			}

			if err = SkillPersistence.Load(path.Join(storagePath, skillPersistenceFileName)); err != nil {
				errChan <- err
			}
		}

		loaded = true
	}
}

// Quit 持久化模块退出前将缓存写入文件
func Quit() {
	RuleSetPersistence.Exit()
	CardPersistence.Exit()
	CharacterPersistence.Exit()
	SkillPersistence.Exit()
}
