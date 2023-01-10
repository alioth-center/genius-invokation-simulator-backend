package persistence

import (
	"sync"
	"time"
)

type timingCacheRecord[T any] struct {
	data      T
	timeoutAt time.Time
}

// timingMemoryCache 带超时系统的的内存缓存Map，实现同kv.syncMap
type timingMemoryCache[PK comparable, T any] struct {
	mutex sync.RWMutex
	cache map[PK]*timingCacheRecord[T]
}

func (t *timingMemoryCache[PK, T]) get(id PK) (exist bool, record *timingCacheRecord[T]) {
	t.mutex.RLock()
	r, has := t.cache[id]
	t.mutex.RUnlock()
	if !has {
		return false, record
	} else if !r.timeoutAt.IsZero() && r.timeoutAt.Before(time.Now()) {
		t.mutex.Lock()
		delete(t.cache, id)
		t.mutex.Unlock()
		return false, record
	} else {
		return true, r
	}
}

func (t *timingMemoryCache[PK, T]) QueryByID(id PK) (has bool, result T, timeoutAt time.Time) {
	if exist, record := t.get(id); !exist {
		return false, result, time.Time{}
	} else {
		return true, record.data, record.timeoutAt
	}
}

func (t *timingMemoryCache[PK, T]) UpdateByID(id PK, entity T) (success bool, timeoutAt time.Time) {
	if exist, record := t.get(id); !exist {
		return false, time.Time{}
	} else {
		record.data = entity
		return true, record.timeoutAt
	}
}

func (t *timingMemoryCache[PK, T]) RefreshByID(id PK, timeout time.Duration) (success bool, timeoutAt time.Time) {
	if exist, record := t.get(id); !exist {
		return false, time.Time{}
	} else if timeout == 0 {
		record.timeoutAt = time.Time{}
		return true, record.timeoutAt
	} else {
		record.timeoutAt = time.Now().Add(timeout)
		return true, record.timeoutAt
	}
}

func (t *timingMemoryCache[PK, T]) InsertOne(id PK, entity T, timeout time.Duration) (success bool, timeoutAt time.Time) {
	if exist, _ := t.get(id); exist {
		return false, time.Time{}
	} else {
		if timeout == 0 {
			t.mutex.Lock()
			t.cache[id] = &timingCacheRecord[T]{data: entity, timeoutAt: time.Time{}}
			t.mutex.Unlock()
			return true, time.Time{}
		} else {
			timeoutAt = time.Now().Add(timeout)
			t.mutex.Lock()
			t.cache[id] = &timingCacheRecord[T]{data: entity, timeoutAt: timeoutAt}
			t.mutex.Unlock()
			return true, timeoutAt
		}
	}
}

func (t *timingMemoryCache[PK, T]) DeleteByID(id PK) (success bool) {
	if exist, _ := t.get(id); exist {
		t.mutex.Lock()
		delete(t.cache, id)
		t.mutex.Unlock()
		return true
	} else {
		return false
	}
}

func newTimingMemoryCache[PK comparable, T any]() TimingMemoryCache[PK, T] {
	return &timingMemoryCache[PK, T]{
		mutex: sync.RWMutex{},
		cache: map[PK]*timingCacheRecord[T]{},
	}
}
