package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type BaseModifier interface {
	ModifierID() uint
	ModifierType() enum.ModifierType
	RoundStartReset()
	Effective() bool
	EffectLeft() uint
}

type AttackModifier interface {
	BaseModifier
	ModifyAttack() func(ctx *modifier.Context[context.DamageContext])
	Clone() AttackModifier
}

type CostModifier interface {
	BaseModifier
	ModifyCost() func(ctx *modifier.Context[context.CostContext])
	Clone() CostModifier
}

type HealModifier interface {
	BaseModifier
	ModifyHeal() func(ctx *modifier.Context[context.HealContext])
	Clone() HealModifier
}

type ChargeModifier interface {
	BaseModifier
	ModifyCharacter() func(ctx *modifier.Context[context.ChargeContext])
	Clone() ChargeModifier
}
