package persistence

import "time"

// FactoryPersistenceRecord 抽象工厂的持久化结构记录
type FactoryPersistenceRecord struct {
	ID  uint   `json:"id"`
	UID string `json:"uid"`
}

// FactoryPersistence 持久化接口，抽象工厂集合的持久化封装
type FactoryPersistence[T any] interface {
	Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error)
	Exit()
	Load(filePath string) (err error)
	QueryByID(id uint) (has bool, result Factory[T])
	QueryByUID(uid string) (has bool, result Factory[T])
	Register(ctor func() T) (success bool)
	Flush(flushPath string, flushFile string) (err error)
}

// Increasable 可增长的，用于数据库主键
type Increasable interface {
	int | uint | byte | rune | int8 | int16 | int64 | uint16 | uint32 | uint64
}

// DatabasePersistence 数据库持久化
type DatabasePersistence[PK Increasable, T any] interface {
	QueryByID(id PK) (has bool, result T)
	UpdateByID(id PK, newEntity T) (success bool)
	InsertOne(entity T) (success bool, id PK)
	DeleteOne(id PK) (success bool)
	FindOne(condition T) (has bool, result T)
	Find(condition T) (results []T)
}

// MemoryCache 内存缓存，不进行持久化
type MemoryCache[PK comparable, T any] interface {
	QueryByID(id PK) (has bool, result T)
	UpdateByID(id PK, newEntity T) (success bool)
	InsertOne(id PK, entity T) (success bool)
	DeleteOne(id PK) (success bool)
}
