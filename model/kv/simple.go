package kv

type Simple[key comparable, value any] struct {
	data map[key]value
}

func (m Simple[key, value]) Exists(key key) bool {
	_, ok := m.data[key]
	return ok
}

func (m Simple[key, value]) Get(key key) value {
	return m.data[key]
}

func (m *Simple[key, value]) Set(key key, data value) {
	m.data[key] = data
}

func (m *Simple[key, value]) Remove(key key) {
	delete(m.data, key)
}

func (m *Simple[key, value]) Range(f func(key, value) bool) {
	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}

func NewSimpleMap[value any]() Map[uint, value] {
	return &Simple[uint, value]{data: map[uint]value{}}
}

func NewCommonMap[key comparable, value any]() Map[key, value] {
	return &Simple[key, value]{data: map[key]value{}}
}
