package kv

import "sync"

type Sync[value any] struct {
	mutex *sync.RWMutex
	data  map[uint]value
}

func (s *Sync[value]) Exists(key uint) bool {
	s.mutex.RLock()
	_, ok := s.data[key]
	s.mutex.RUnlock()
	return ok
}

func (s *Sync[value]) Get(key uint) value {
	s.mutex.RLock()
	result := s.data[key]
	s.mutex.RUnlock()
	return result
}

func (s *Sync[value]) Set(key uint, data value) {
	s.mutex.Lock()
	s.data[key] = data
	s.mutex.Unlock()
}

func (s *Sync[value]) Remove(key uint) {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
}

func (s *Sync[value]) Range(f func(uint, value) bool) {
	s.mutex.RLock()
	for k, v := range s.data {
		if !f(k, v) {
			break
		}
	}
	s.mutex.RUnlock()
}

func NewSyncMap[value any]() Map[uint, value] {
	return &Sync[value]{
		mutex: &sync.RWMutex{},
		data:  map[uint]value{},
	}
}
