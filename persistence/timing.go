package persistence

import (
	"math/rand"
	"sync"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type timingCacheRecord[T any] struct {
	data      T
	timeoutAt time.Time
}

// timingMemoryCache 带超时系统的的内存缓存Map，实现同kv.syncMap
type timingMemoryCache[PK comparable, T any] struct {
	mutex  sync.RWMutex
	cache  map[PK]*timingCacheRecord[T]
	size   uint
	exit   chan struct{}
	served bool
}

func (t *timingMemoryCache[PK, T]) proactivelyClean(index float64) {
	if index < 0 || index > 1 {
		index = 1
	}

	need, executeCount := uint(float64(t.size)/index), uint(0)
	executeList := make([]PK, need)

	// 统计需要删除的超时记录
	t.mutex.RLock()
	for pk, entity := range t.cache {
		if !entity.timeoutAt.IsZero() && entity.timeoutAt.Before(time.Now()) {
			executeList[executeCount] = pk
			executeCount++
		}

		if need = need - 1; need == 0 {
			break
		}
	}
	t.mutex.RUnlock()

	// 执行删除操作，频繁加锁性能较差，此处不采用分段锁
	t.mutex.Lock()
	for i := uint(0); i < executeCount; i++ {
		delete(t.cache, executeList[i])
		t.size -= 1
	}
	t.mutex.Unlock()
}

func (t *timingMemoryCache[PK, T]) serve(proactivelyCleanTime time.Duration, proactivelyCleanIndex float64) {
	// 起服务前先随机sleep一会，让一起初始化的多个组建主动清理时间错开
	randomNum := random.Int63n(int64(proactivelyCleanTime))
	time.Sleep(time.Duration(randomNum))

	for {
		select {
		case <-t.exit:
			return
		default:
			t.proactivelyClean(proactivelyCleanIndex)
			time.Sleep(proactivelyCleanTime)
		}
	}
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
		t.size -= 1
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
			t.size += 1
			t.mutex.Unlock()
			return true, time.Time{}
		} else {
			timeoutAt = time.Now().Add(timeout)
			t.mutex.Lock()
			t.cache[id] = &timingCacheRecord[T]{data: entity, timeoutAt: timeoutAt}
			t.size += 1
			t.mutex.Unlock()
			return true, timeoutAt
		}
	}
}

func (t *timingMemoryCache[PK, T]) DeleteByID(id PK) (success bool) {
	if exist, _ := t.get(id); exist {
		t.mutex.Lock()
		delete(t.cache, id)
		t.size -= 1
		t.mutex.Unlock()
		return true
	} else {
		return false
	}
}

func (t *timingMemoryCache[PK, T]) Serve(proactivelyCleanTime time.Duration, proactivelyCleanIndex float64) {
	if proactivelyCleanTime > 0 {
		go t.serve(proactivelyCleanTime, proactivelyCleanIndex)
		t.served = true
	} else {
		t.served = false
	}
}

func (t *timingMemoryCache[PK, T]) Exit() {
	if t.served {
		t.exit <- struct{}{}
	}
}

func newTimingMemoryCache[PK comparable, T any]() TimingMemoryCache[PK, T] {
	return &timingMemoryCache[PK, T]{
		mutex:  sync.RWMutex{},
		cache:  map[PK]*timingCacheRecord[T]{},
		size:   0,
		exit:   make(chan struct{}, 1),
		served: false,
	}
}
