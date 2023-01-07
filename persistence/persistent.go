package persistence

// Persistent 可持久化的实体接口，本质是一个生产entity的工厂
type Persistent[entity any] interface {
	Ctor() func() entity
	Enable() bool
	ID() uint
	UID() string

	set(id uint, uid string, ctor func() entity)
	enable()
	disable()
}

func NewPersistent[entity any](id uint, uid string) Persistent[entity] {
	return &persistent[entity]{
		id:     id,
		uid:    uid,
		status: false,
		ctor:   nil,
	}
}

type persistent[entity any] struct {
	ctor   func() entity
	status bool
	id     uint
	uid    string
}

func (p *persistent[entity]) Ctor() func() entity {
	return p.ctor
}

func (p *persistent[entity]) Enable() bool {
	return p.status
}

func (p *persistent[entity]) ID() uint {
	return p.id
}

func (p *persistent[entity]) UID() string {
	return p.uid
}

func (p *persistent[entity]) set(id uint, uid string, ctor func() entity) {
	p.id, p.uid, p.ctor = id, uid, ctor
}

func (p *persistent[entity]) disable() {
	p.status = false
}

func (p *persistent[entity]) enable() {
	p.status = true
}
