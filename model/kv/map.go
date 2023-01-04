package kv

// Pair key-value存储对的封装
type Pair[key, value any] interface {
	Key() key
	Value() value
	SetKey(key)
	SetValue(value)
}

// Map key-value存储的封装
type Map[key comparable, value any] interface {
	Exists(key) bool
	Get(key) value
	Set(key, value)
	Remove(key)
	Range(func(key, value) bool)
}

// OrderedMap 有序Map的封装
type OrderedMap[key comparable, value any] interface {
	Map[key, value]
	Length() uint
	GetIndex(key) uint
	GetKey(uint) key
}
