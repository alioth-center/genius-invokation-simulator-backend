package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type AttackModifiers = modifier.Chain[context.DamageContext]

type DefenceModifiers = modifier.Chain[context.DamageContext]

type HealModifiers = modifier.Chain[context.HealContext]

type ChargeModifiers = modifier.Chain[context.ChargeContext]

type CostModifiers = modifier.Chain[context.CostContext]

type AttackModifier = modifier.Modifier[context.DamageContext]

type DefenceModifier = modifier.Modifier[context.DamageContext]

type HealModifier = modifier.Modifier[context.HealContext]

type ChargeModifier = modifier.Modifier[context.ChargeContext]

type CostModifier = modifier.Modifier[context.CostContext]
