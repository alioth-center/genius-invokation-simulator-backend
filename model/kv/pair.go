package kv

type pair[key, value any] struct {
	k key
	v value
}

func (p pair[key, value]) Key() key {
	return p.k
}

func (p pair[key, value]) Value() value {
	return p.v
}

func (p *pair[key, value]) SetKey(k key) {
	p.k = k
}

func (p *pair[key, value]) SetValue(v value) {
	p.v = v
}

func NewPair[key, value any](k key, v value) Pair[key, value] {
	return &pair[key, value]{
		k: k,
		v: v,
	}
}
