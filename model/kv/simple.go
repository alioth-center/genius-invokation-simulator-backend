package kv

type Simple[key comparable, value any] struct {
	data map[key]value
}

func (m Simple[key, value]) Exists(k key) bool {
	_, ok := m.data[k]
	return ok
}

func (m Simple[key, value]) Get(k key) value {
	return m.data[k]
}

func (m *Simple[key, value]) Set(k key, data value) {
	m.data[k] = data
}

func (m *Simple[key, value]) Remove(k key) {
	delete(m.data, k)
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
