package modifier

type Modifier[data any] interface {
	Info() Info // Info 基础信息
	Handler() func(ctx *Context[data])
	Clone() Modifier[data]
	RoundReset()
	Effected() bool
	EffectLeft() uint
}

type Info struct {
	ID          uint
	Name        string
	Description string
}
