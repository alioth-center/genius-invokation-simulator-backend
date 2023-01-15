package modifier

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Modifier[data any] interface {
	ID() uint
	Type() enum.ModifierType
	Handler() func(ctx *Context[data])
	Clone() Modifier[data]
	RoundReset()
	Effective() bool
	EffectLeft() uint
}
