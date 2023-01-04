package kv

type simpleMap[key comparable, value any] struct {
	data map[key]value
}

func (m simpleMap[key, value]) Exists(k key) bool {
	_, ok := m.data[k]
	return ok
}

func (m simpleMap[key, value]) Get(k key) value {
	return m.data[k]
}

func (m *simpleMap[key, value]) Set(k key, data value) {
	m.data[k] = data
}

func (m *simpleMap[key, value]) Remove(k key) {
	delete(m.data, k)
}

func (m *simpleMap[key, value]) Range(f func(key, value) bool) {
	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}

func NewSimpleMap[value any]() Map[uint, value] {
	return &simpleMap[uint, value]{data: map[uint]value{}}
}

func NewCommonMap[key comparable, value any]() Map[key, value] {
	return &simpleMap[key, value]{data: map[key]value{}}
}
