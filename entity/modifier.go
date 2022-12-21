package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type AttackModifiers = modifier.Chain[context.DamageContext]

type DefenceModifiers = modifier.Chain[context.DamageContext]

type HealModifiers = modifier.Chain[context.HealContext]

type ChargeModifiers = modifier.Chain[context.ChargeContext]

type CostModifiers = modifier.Chain[context.CostContext]
