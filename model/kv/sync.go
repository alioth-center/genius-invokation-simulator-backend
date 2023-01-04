package kv

import "sync"

type syncMap[value any] struct {
	mutex *sync.RWMutex
	data  map[uint]value
}

func (s *syncMap[value]) Exists(key uint) bool {
	s.mutex.RLock()
	_, ok := s.data[key]
	s.mutex.RUnlock()
	return ok
}

func (s *syncMap[value]) Get(key uint) value {
	s.mutex.RLock()
	result := s.data[key]
	s.mutex.RUnlock()
	return result
}

func (s *syncMap[value]) Set(key uint, data value) {
	s.mutex.Lock()
	s.data[key] = data
	s.mutex.Unlock()
}

func (s *syncMap[value]) Remove(key uint) {
	s.mutex.Lock()
	delete(s.data, key)
	s.mutex.Unlock()
}

func (s *syncMap[value]) Range(f func(uint, value) bool) {
	s.mutex.RLock()
	for k, v := range s.data {
		if !f(k, v) {
			break
		}
	}
	s.mutex.RUnlock()
}

func NewSyncMap[value any]() Map[uint, value] {
	return &syncMap[value]{
		mutex: &sync.RWMutex{},
		data:  map[uint]value{},
	}
}
