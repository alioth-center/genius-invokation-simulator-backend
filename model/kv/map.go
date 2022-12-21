package kv

// Map key-value存储的封装
type Map[key comparable, value any] interface {
	Exists(key) bool
	Get(key) value
	Set(key, value)
	Remove(key)
	Range(func(key, value) bool)
}
