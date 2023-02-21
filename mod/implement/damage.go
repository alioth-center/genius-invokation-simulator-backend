package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type DamageImpl struct {
	elementType  enum.ElementType
	damageAmount uint
}

func (impl *DamageImpl) ElementType() enum.ElementType {
	return impl.elementType
}

func (impl *DamageImpl) DamageAmount() uint {
	return impl.damageAmount
}

type DamageOptions func(option *DamageImpl)

func WithDamageElementType(elementType enum.ElementType) DamageOptions {
	return func(option *DamageImpl) {
		option.elementType = elementType
	}
}

func WithDamageAmount(amount uint) DamageOptions {
	return func(option *DamageImpl) {
		option.damageAmount = amount
	}
}

func NewDamageWithOpts(options ...DamageOptions) definition.Damage {
	impl := &DamageImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}
