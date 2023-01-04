package kv

type orderedMap[key comparable, value any] struct {
	data   map[key]value
	order  map[uint]key
	cache  map[key]uint
	length uint
}

func (o orderedMap[key, value]) Exists(k key) bool {
	_, ok := o.data[k]
	return ok
}

func (o orderedMap[key, value]) Get(k key) value {
	return o.data[k]
}

func (o *orderedMap[key, value]) Set(k key, v value) {
	if o.Exists(k) {
		o.data[k] = v
	} else {
		o.order[o.length] = k
		o.cache[k] = o.length
		o.length += 1
		o.data[k] = v
	}
}

func (o *orderedMap[key, value]) Remove(k key) {
	if o.Exists(k) {
		for i := o.GetIndex(k) + 1; i < o.length; i++ {
			o.order[i-1] = o.order[i]
			o.cache[o.order[i-1]] = i - 1
		}
		o.length -= 1
		delete(o.data, k)
		delete(o.cache, k)
		delete(o.order, o.length)
	}
}

func (o orderedMap[key, value]) Length() uint {
	return o.length
}

func (o orderedMap[key, value]) GetIndex(k key) uint {
	return o.cache[k]
}

func (o orderedMap[key, value]) GetKey(index uint) key {
	return o.order[index]
}

func (o orderedMap[key, value]) Range(f func(key, value) bool) {
	for i := uint(0); i < o.length; i++ {
		if !f(o.order[i], o.data[o.order[i]]) {
			break
		}
	}
}

func NewOrderedMap[key comparable, value any]() OrderedMap[key, value] {
	return &orderedMap[key, value]{
		data:   map[key]value{},
		order:  map[uint]key{},
		cache:  map[key]uint{},
		length: 0,
	}
}
