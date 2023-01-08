package persistence

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
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

func init() {
	RuleSetPersistence = NewPersistence[RuleSet]()
	CardPersistence = NewPersistence[Card]()
	CharacterPersistence = NewPersistence[Character]()
	SkillPersistence = NewPersistence[Skill]()
}

var (
	RuleSetPersistence   Persistence[RuleSet]
	CardPersistence      Persistence[Card]
	CardDeckPersistence  Persistence[CardDeck]
	CharacterPersistence Persistence[Character]
	PlayerPersistence    Persistence[Player]
	SkillPersistence     Persistence[Skill]
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
	PlayerPersistence.Serve(flushFeq, storagePath, sqlite3DBFileName, errChan)
	CardDeckPersistence.Serve(flushFeq, storagePath, sqlite3DBFileName, errChan)
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

			CardDeckPersistence = NewSqlite3Table[CardDeck]()
			PlayerPersistence = NewSqlite3Table[Player]()
		}

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

// Persistence 持久化接口，抽象工厂集合的持久化封装
type Persistence[T any] interface {
	Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error)
	Exit()
	Load(filePath string) (err error)
	QueryByID(id uint) (has bool, result Persistent[T])
	QueryByUID(uid string) (has bool, result Persistent[T])
	Register(ctor func() T) (success bool)
	Flush(flushPath string, flushFile string) (err error)
}

// persistence 持久化接口的实现
type persistence[T any] struct {
	impl *PerformanceMap[T]
	exit chan struct{}
}

func (p *persistence[T]) Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error) {
	go func() {
		// 执行定时任务
		exitCh := make(chan struct{})
		go func(ch chan struct{}) {
			for {
				select {
				case <-exitCh:
					return
				default:
					// 每隔flushFrequency，将缓存写入文件
					if err := p.Flush(flushPath, flushFile); err != nil {
						errChan <- err
					}
					time.Sleep(flushFrequency)
				}
			}
		}(exitCh)

		// 监听exit信号
		<-p.exit
		exitCh <- struct{}{}
		// 收到退出信号，立即将缓存写入文件
		if err := p.Flush(flushPath, strings.Join([]string{flushFile, "quit"}, ".")); err != nil {
			errChan <- err
		}
	}()
}

func (p *persistence[T]) Exit() {
	p.exit <- struct{}{}
}

func (p *persistence[T]) Load(filePath string) (err error) {
	if file, err := os.Open(filePath); err != nil {
		return err
	} else if fileContent, err := io.ReadAll(file); err != nil {
		return err
	} else {
		var persistenceEntities []PerformanceMapRecord
		if err = json.Unmarshal(fileContent, &persistenceEntities); err != nil {
			return err
		} else {
			p.impl.Load(persistenceEntities)
			return nil
		}
	}
}

func (p *persistence[T]) QueryByID(id uint) (has bool, result Persistent[T]) {
	return p.impl.QueryByID(id)
}

func (p *persistence[T]) QueryByUID(uid string) (has bool, result Persistent[T]) {
	return p.impl.QueryByUID(uid)
}

func (p *persistence[T]) Register(ctor func() T) (success bool) {
	return p.impl.Register(ctor)
}

func (p *persistence[T]) Flush(flushPath string, flushFile string) (err error) {
	persistenceEntities := p.impl.Flush()
	if fileContent, err := json.Marshal(&persistenceEntities); err != nil {
		return err
	} else if file, err := os.OpenFile(path.Join(flushPath, flushFile), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {
		return err
	} else {
		_, err = file.Write(fileContent)
		return err
	}
}

func NewPersistence[T any]() Persistence[T] {
	return NewPersistenceWithHighPerformance[T]()
}

func NewPersistenceWithImpl[T any](impl *PerformanceMap[T]) Persistence[T] {
	return &persistence[T]{
		impl: impl,
		exit: make(chan struct{}, 1),
	}
}

func NewPersistenceWithLowPerformance[T any]() Persistence[T] {
	_, impl := NewPerformanceMapWithOpts[T](8)
	return &persistence[T]{
		impl: impl,
		exit: make(chan struct{}, 1),
	}
}

func NewPersistenceWithMediumPerformance[T any]() Persistence[T] {
	_, impl := NewPerformanceMapWithOpts[T](64)
	return &persistence[T]{
		impl: impl,
		exit: make(chan struct{}, 1),
	}
}

func NewPersistenceWithHighPerformance[T any]() Persistence[T] {
	return &persistence[T]{
		impl: NewPerformanceMap[T](),
		exit: make(chan struct{}, 1),
	}
}
