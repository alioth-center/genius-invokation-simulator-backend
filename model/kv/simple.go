package kv

type Simple[value any] struct {
	data map[uint]value
}

func (m Simple[value]) Exists(key uint) bool {
	_, ok := m.data[key]
	return ok
}

func (m Simple[value]) Get(key uint) value {
	return m.data[key]
}

func (m *Simple[value]) Set(key uint, data value) {
	m.data[key] = data
}

func (m *Simple[value]) Remove(key uint) {
	delete(m.data, key)
}

func (m *Simple[value]) Range(f func(uint, value) bool) {
	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}

func NewSimpleMap[value any]() Map[uint, value] {
	return &Simple[value]{data: map[uint]value{}}
}
