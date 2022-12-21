package modifier

type Modifier[data any] interface {
	ID() uint
	Handler() func(ctx *Context[data])
	Clone() Modifier[data]
	RoundReset()
	Effective() bool
	EffectLeft() uint
}
