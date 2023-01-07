package persistence

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

const (
	ruleSetPersistenceFileName   = "rule-set-persistence.json"
	cardPersistenceFileName      = "card-persistence.json"
	characterPersistenceFileName = "character-persistence.json"
	playerPersistenceFileName    = "player-persistence.json"
	skillPersistenceFileName     = "skill-persistence.json"
)

var (
	storagePath = ""
	loaded      = false
)

var (
// RuleSetPersistence   Persistence[RuleSet]
// CardPersistence      Persistence[Card]
// CharacterPersistence Persistence[CharacterInfo]
// PlayerPersistence    Persistence[PlayerInfo]
// SkillPersistence     Persistence[Skill]
)

func init() {
	if execPath, err := os.Executable(); err != nil {
		panic(err)
	} else if err = SetStoragePath(path.Join(execPath, "../data/persistence")); err != nil {
		panic(err)
	} else {
		fmt.Println(path.Join(storagePath, ruleSetPersistenceFileName))
	}
}

func SetStoragePath(path string) error {
	if s, err := os.Stat(path); err == nil && s.IsDir() {
		storagePath = path
		return nil
	} else {
		return fmt.Errorf("path %s is not a directory or is not exist", path)
	}
}

// Load 从持久化文件读取信息，写入持久化模块
func Load() {
	if !loaded {
		//RuleSetPersistence.Load(path.Join(storagePath, ruleSetPersistenceFileName))
		//CardPersistence.Load(path.Join(storagePath, cardPersistenceFileName))
		//CharacterPersistence.Load(path.Join(storagePath, characterPersistenceFileName))
		//PlayerPersistence.Load(path.Join(storagePath, playerPersistenceFileName))
		//SkillPersistence.Load(path.Join(storagePath, skillPersistenceFileName))
		loaded = true
	}
}

// Quit 持久化模块退出前将缓存写入文件
func Quit() {
	//RuleSetPersistence.Flush(storagePath, ruleSetPersistenceFileName)
	//CardPersistence.Flush(storagePath, cardPersistenceFileName)
	//CharacterPersistence.Flush(storagePath, characterPersistenceFileName)
	//PlayerPersistence.Flush(storagePath, playerPersistenceFileName)
	//SkillPersistence.Flush(storagePath, skillPersistenceFileName)
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
		for {
			select {
			case <-p.exit:
				// 收到退出信号，立即将缓存写入文件
				if err := p.Flush(flushPath, flushFile); err != nil {
					errChan <- err
				}
				return
			default:
				// 每隔flushFrequency，将缓存写入文件
				time.Sleep(flushFrequency)
				if err := p.Flush(flushPath, flushFile); err != nil {
					errChan <- err
				}
			}
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
		exit: make(chan struct{}),
	}
}

func NewPersistenceWithLowPerformance[T any]() Persistence[T] {
	_, impl := NewPerformanceMapWithOpts[T](8)
	return &persistence[T]{
		impl: impl,
		exit: make(chan struct{}),
	}
}

func NewPersistenceWithMediumPerformance[T any]() Persistence[T] {
	_, impl := NewPerformanceMapWithOpts[T](64)
	return &persistence[T]{
		impl: impl,
		exit: make(chan struct{}),
	}
}

func NewPersistenceWithHighPerformance[T any]() Persistence[T] {
	return &persistence[T]{
		impl: NewPerformanceMap[T](),
		exit: make(chan struct{}),
	}
}
