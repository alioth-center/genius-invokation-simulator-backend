package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
	model "github.com/sunist-c/genius-invokation-simulator-backend/model/modifier/definition"
)

type baseModifierAdapterLayer struct {
	modifierID   uint64
	modifierType enum.ModifierType
	effective    func() bool
	effectLeft   func() uint
}

func (b *baseModifierAdapterLayer) ID() uint64 {
	return b.modifierID
}

func (b *baseModifierAdapterLayer) Type() enum.ModifierType {
	return b.modifierType
}

func (b *baseModifierAdapterLayer) Effective() bool {
	return b.effective()
}

func (b *baseModifierAdapterLayer) EffectLeft() uint {
	return b.effectLeft()
}

type attackModifierAdapterLayer struct {
	baseModifierAdapterLayer
	attackHandler func(ctx *modifier.Context[context.DamageContext])
	clone         func() definition.AttackModifier
	reset         func()
}

func (a *attackModifierAdapterLayer) Handler() func(ctx *modifier.Context[context.DamageContext]) {
	return a.attackHandler
}

func (a *attackModifierAdapterLayer) Clone() modifier.Modifier[context.DamageContext] {
	clonedSource := a.clone()
	_, clonedResult := attackModifierAdapter.Convert(clonedSource)
	return clonedResult
}

func (a *attackModifierAdapterLayer) RoundReset() {
	a.reset()
}

type AttackModifierAdapter struct{}

func (a AttackModifierAdapter) Convert(source definition.AttackModifier) (success bool, result model.AttackModifier) {
	adapterLayer := &attackModifierAdapterLayer{
		baseModifierAdapterLayer: baseModifierAdapterLayer{
			modifierID:   source.ModifierID(),
			modifierType: source.ModifierType(),
			effective:    source.Effective,
			effectLeft:   source.EffectLeft,
		},
		clone:         source.Clone,
		attackHandler: source.ModifyAttack(),
		reset:         source.RoundStartReset,
	}

	return true, adapterLayer
}

func NewAttackModifierAdapter() adapter.Adapter[definition.AttackModifier, model.AttackModifier] {
	return AttackModifierAdapter{}
}

type costModifierAdapterLayer struct {
	baseModifierAdapterLayer
	costHandler func(ctx *modifier.Context[context.CostContext])
	clone       func() definition.CostModifier
	reset       func()
}

func (c *costModifierAdapterLayer) Handler() func(ctx *modifier.Context[context.CostContext]) {
	return c.costHandler
}

func (c *costModifierAdapterLayer) Clone() modifier.Modifier[context.CostContext] {
	clonedSource := c.clone()
	_, clonedResult := costModifierAdapter.Convert(clonedSource)
	return clonedResult
}

func (c *costModifierAdapterLayer) RoundReset() {
	c.reset()
}

type CostModifierAdapter struct{}

func (c CostModifierAdapter) Convert(source definition.CostModifier) (success bool, result model.CostModifier) {
	adapterLayer := &costModifierAdapterLayer{
		baseModifierAdapterLayer: baseModifierAdapterLayer{
			modifierID:   source.ModifierID(),
			modifierType: source.ModifierType(),
			effective:    source.Effective,
			effectLeft:   source.EffectLeft,
		},
		costHandler: source.ModifyCost(),
		clone:       source.Clone,
		reset:       source.RoundStartReset,
	}

	return true, adapterLayer
}

func NewCostModifierAdapter() adapter.Adapter[definition.CostModifier, model.CostModifier] {
	return CostModifierAdapter{}
}

type healModifierAdapterLayer struct {
	baseModifierAdapterLayer
	healHandler func(ctx *modifier.Context[context.HealContext])
	clone       func() definition.HealModifier
	reset       func()
}

func (h *healModifierAdapterLayer) Handler() func(ctx *modifier.Context[context.HealContext]) {
	return h.healHandler
}

func (h *healModifierAdapterLayer) Clone() modifier.Modifier[context.HealContext] {
	clonedSource := h.clone()
	_, clonedResult := healModifierAdapter.Convert(clonedSource)
	return clonedResult
}

func (h *healModifierAdapterLayer) RoundReset() {
	h.reset()
}

type HealModifierAdapter struct{}

func (h HealModifierAdapter) Convert(source definition.HealModifier) (success bool, result model.HealModifier) {
	adapterLayer := &healModifierAdapterLayer{
		baseModifierAdapterLayer: baseModifierAdapterLayer{
			modifierID:   source.ModifierID(),
			modifierType: source.ModifierType(),
			effective:    source.Effective,
			effectLeft:   source.EffectLeft,
		},
		healHandler: source.ModifyHeal(),
		clone:       source.Clone,
		reset:       source.RoundStartReset,
	}

	return true, adapterLayer
}

func NewHealModifierAdapter() adapter.Adapter[definition.HealModifier, model.HealModifier] {
	return HealModifierAdapter{}
}

type chargeModifierAdapterLayer struct {
	baseModifierAdapterLayer
	chargeHandler func(ctx *modifier.Context[context.ChargeContext])
	clone         func() definition.ChargeModifier
	reset         func()
}

func (c *chargeModifierAdapterLayer) Handler() func(ctx *modifier.Context[context.ChargeContext]) {
	return c.chargeHandler
}

func (c *chargeModifierAdapterLayer) Clone() modifier.Modifier[context.ChargeContext] {
	clonedSource := c.clone()
	_, clonedResult := chargeModifierAdapter.Convert(clonedSource)
	return clonedResult
}

func (c *chargeModifierAdapterLayer) RoundReset() {
	c.reset()
}

type ChargeModifierAdapter struct{}

func (c ChargeModifierAdapter) Convert(source definition.ChargeModifier) (success bool, result model.ChargeModifier) {
	adapterLayer := &chargeModifierAdapterLayer{
		baseModifierAdapterLayer: baseModifierAdapterLayer{
			modifierID:   source.ModifierID(),
			modifierType: source.ModifierType(),
			effective:    source.Effective,
			effectLeft:   source.EffectLeft,
		},
		chargeHandler: source.ModifyCharacter(),
		clone:         source.Clone,
		reset:         source.RoundStartReset,
	}

	return true, adapterLayer
}

func NewChargeModifierAdapter() adapter.Adapter[definition.ChargeModifier, model.ChargeModifier] {
	return ChargeModifierAdapter{}
}
