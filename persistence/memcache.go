package persistence

import (
	"sync"
)

// memoryCache 内存缓存的实现，简单的并发安全map封装，同kv.syncMap实现
type memoryCache[PK comparable, T any] struct {
	mu    sync.RWMutex
	cache map[PK]T
}

func (m *memoryCache[PK, T]) QueryByID(id PK) (has bool, result T) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if entity, exist := m.cache[id]; !exist {
		return false, result
	} else {
		return true, entity
	}
}

func (m *memoryCache[PK, T]) UpdateByID(id PK, newEntity T) (success bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if _, exist := m.cache[id]; !exist {
		return false
	} else {
		m.cache[id] = newEntity
		return true
	}
}

func (m *memoryCache[PK, T]) InsertOne(id PK, entity T) (success bool) {
	m.mu.RLock()
	_, exist := m.cache[id]
	m.mu.RUnlock()

	if !exist {
		m.mu.Lock()
		m.cache[id] = entity
		m.mu.Unlock()
		return true
	} else {
		return false
	}
}

func (m *memoryCache[PK, T]) DeleteOne(id PK) (success bool) {
	m.mu.RLock()
	_, exist := m.cache[id]
	m.mu.RUnlock()

	if !exist {
		return false
	} else {
		m.mu.Lock()
		delete(m.cache, id)
		m.mu.Unlock()
		return true
	}
}

func newMemoryCache[PK comparable, T any]() MemoryCache[PK, T] {
	return &memoryCache[PK, T]{
		mu:    sync.RWMutex{},
		cache: map[PK]T{},
	}
}
