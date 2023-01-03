package kv

type ordered[key comparable, value any] struct {
	data   map[key]value
	order  map[uint]key
	cache  map[key]uint
	length uint
}

func (o ordered[key, value]) Exists(k key) bool {
	_, ok := o.data[k]
	return ok
}

func (o ordered[key, value]) Get(k key) value {
	return o.data[k]
}

func (o *ordered[key, value]) Set(k key, v value) {
	if o.Exists(k) {
		o.data[k] = v
	} else {
		o.order[o.length] = k
		o.cache[k] = o.length
		o.length += 1
		o.data[k] = v
	}
}

func (o *ordered[key, value]) Remove(k key) {
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

func (o ordered[key, value]) Length() uint {
	return o.length
}

func (o ordered[key, value]) GetIndex(k key) uint {
	return o.cache[k]
}

func (o ordered[key, value]) GetKey(index uint) key {
	return o.order[index]
}

func (o ordered[key, value]) Range(start, end uint, f func(key, value)) {
	for i := start; i < end; i++ {
		f(o.order[i], o.data[o.order[i]])
	}
}

func NewOrderedMap[key comparable, value any]() OrderedMap[key, value] {
	return &ordered[key, value]{
		data:   map[key]value{},
		order:  map[uint]key{},
		cache:  map[key]uint{},
		length: 0,
	}
}
