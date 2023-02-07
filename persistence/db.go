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
	summonPersistenceFileName    = "summon-persistence.psc"
	eventPersistenceFileName     = "event-persistence.psc"
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
	SummonPersistence    = newFactoryPersistence[Summon]()
	EventPersistence     = newFactoryPersistence[Event]()

	LocalizationPersistence = newMemoryCache[string, localization.LanguagePack]()
	ModInfoPersistence      = newMemoryCache[string, ModInfo]()
	RoomInfoPersistence     = newMemoryCache[uint64, RoomInfo]()

	TokenPersistence = newTimingMemoryCache[string, Token]()

	CardDeckPersistence DatabasePersistence[uint64, CardDeck]
	PlayerPersistence   DatabasePersistence[uint64, Player]
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
	SummonPersistence.Serve(flushFeq, storagePath, summonPersistenceFileName, errChan)
	EventPersistence.Serve(flushFeq, storagePath, eventPersistenceFileName, errChan)
	TokenPersistence.Serve(time.Second*time.Duration(300), 0.5)
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

			var success bool
			if success, CardDeckPersistence = newDatabasePersistence[uint64, CardDeck](errChan); !success {
				errChan <- fmt.Errorf("failed to create database factoryPersistence with entity: %+v", CardDeck{})
			}

			if success, PlayerPersistence = newDatabasePersistence[uint64, Player](errChan); !success {
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

			if err = SummonPersistence.Load(path.Join(storagePath, summonPersistenceFileName)); err != nil {
				errChan <- err
			}

			if err = EventPersistence.Load(path.Join(storagePath, eventPersistenceFileName)); err != nil {
				errChan <- err
			}
		}

		loaded = true
	}
}

// Quit 退出持久化模块的各种持久化服务
func Quit() {
	RuleSetPersistence.Exit()
	CardPersistence.Exit()
	CharacterPersistence.Exit()
	SkillPersistence.Exit()
	SummonPersistence.Exit()
	EventPersistence.Exit()
	TokenPersistence.Exit()
}
