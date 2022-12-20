package modifier

type Modifier[data any] interface {
	Info() Info
	Handler() func(ctx *Context[data])
	Clone() Modifier[data]
	RoundReset()
	Effective() bool
	EffectLeft() uint
}

// Info Modifier的信息
type Info struct {
	id          uint
	name        string
	description string
}

// ID Modifier的ID
func (i Info) ID() uint { return i.id }

// Name Modifier的名称
func (i Info) Name() string { return i.name }

// Description Modifier的描述
func (i Info) Description() string { return i.description }
