package persistence

import (
	"fmt"
	"os"
	"path"
)

var (
	storagePath = ""
)

func init() {
	if execPath, err := os.Executable(); err != nil {
		panic(err)
	} else if err = SetStoragePath(path.Join(execPath, "../data/persistence")); err != nil {
		panic(err)
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

var (
	RuleSetPersistence   Persistence[RuleSet]
	CardPersistence      Persistence[Card]
	CharacterPersistence Persistence[CharacterInfo]
	PlayerPersistence    Persistence[PlayerInfo]
	SkillPersistence     Persistence[Skill]
)

// PersistenceIndex 持久化索引，本质上是一个抽象工厂
type PersistenceIndex[T any] struct {
	ID   uint
	UID  string
	Ctor func() T
}

// Persistence 持久化接口，抽象工厂集合的持久化封装
type Persistence[T any] interface {
	QueryByID(id uint) (has bool, result T)
	Register(id uint, entity T)
}
