package context

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type ModifierContext struct {
	addLocalDirectAttackModifiers kv.Map[uint, []modifier.Modifier[DamageContext]]
	addLocalFinalAttackModifiers  kv.Map[uint, []modifier.Modifier[DamageContext]]
	addLocalDefenceModifiers      kv.Map[uint, []modifier.Modifier[DamageContext]]
	addLocalChargeModifiers       kv.Map[uint, []modifier.Modifier[ChargeContext]]
	addLocalHealModifiers         kv.Map[uint, []modifier.Modifier[HealContext]]
	addLocalCostModifiers         kv.Map[uint, []modifier.Modifier[CostContext]]

	addGlobalDirectAttackModifiers []modifier.Modifier[DamageContext]
	addGlobalFinalAttackModifiers  []modifier.Modifier[DamageContext]
	addGlobalDefenceModifiers      []modifier.Modifier[DamageContext]
	addGlobalChargeModifiers       []modifier.Modifier[ChargeContext]
	addGlobalHealModifiers         []modifier.Modifier[HealContext]
	addGlobalCostModifiers         []modifier.Modifier[CostContext]

	removeLocalDirectAttackModifiers kv.Map[uint, []modifier.Modifier[DamageContext]]
	removeLocalFinalAttackModifiers  kv.Map[uint, []modifier.Modifier[DamageContext]]
	removeLocalDefenceModifiers      kv.Map[uint, []modifier.Modifier[DamageContext]]
	removeLocalChargeModifiers       kv.Map[uint, []modifier.Modifier[ChargeContext]]
	removeLocalHealModifiers         kv.Map[uint, []modifier.Modifier[HealContext]]
	removeLocalCostModifiers         kv.Map[uint, []modifier.Modifier[CostContext]]

	removeGlobalDirectAttackModifiers []modifier.Modifier[DamageContext]
	removeGlobalFinalAttackModifiers  []modifier.Modifier[DamageContext]
	removeGlobalDefenceModifiers      []modifier.Modifier[DamageContext]
	removeGlobalChargeModifiers       []modifier.Modifier[ChargeContext]
	removeGlobalHealModifiers         []modifier.Modifier[HealContext]
	removeGlobalCostModifiers         []modifier.Modifier[CostContext]
}

func (m *ModifierContext) AddLocalDirectAttackModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.addLocalDirectAttackModifiers != null {
		modifiers := append(m.addLocalDirectAttackModifiers.Get(targetCharacter), handler)
		m.addLocalDirectAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalDirectAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalDirectAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalFinalAttackModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.addLocalFinalAttackModifiers != null {
		modifiers := append(m.addLocalFinalAttackModifiers.Get(targetCharacter), handler)
		m.addLocalFinalAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalFinalAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalFinalAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalDefenceModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.addLocalDefenceModifiers != null {
		modifiers := append(m.addLocalDefenceModifiers.Get(targetCharacter), handler)
		m.addLocalDefenceModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalDefenceModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalDefenceModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalChargeModifier(targetCharacter uint, handler modifier.Modifier[ChargeContext]) {
	var null kv.Map[uint, []modifier.Modifier[ChargeContext]] = nil
	if m.addLocalChargeModifiers != null {
		modifiers := append(m.addLocalChargeModifiers.Get(targetCharacter), handler)
		m.addLocalChargeModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalChargeModifiers = kv.NewSimpleMap[[]modifier.Modifier[ChargeContext]]()
		list := []modifier.Modifier[ChargeContext]{handler}
		m.addLocalChargeModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalHealModifier(targetCharacter uint, handler modifier.Modifier[HealContext]) {
	var null kv.Map[uint, []modifier.Modifier[HealContext]] = nil
	if m.addLocalHealModifiers != null {
		modifiers := append(m.addLocalHealModifiers.Get(targetCharacter), handler)
		m.addLocalHealModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalHealModifiers = kv.NewSimpleMap[[]modifier.Modifier[HealContext]]()
		list := []modifier.Modifier[HealContext]{handler}
		m.addLocalHealModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalCostModifier(targetCharacter uint, handler modifier.Modifier[CostContext]) {
	var null kv.Map[uint, []modifier.Modifier[CostContext]] = nil
	if m.addLocalCostModifiers != null {
		modifiers := append(m.addLocalCostModifiers.Get(targetCharacter), handler)
		m.addLocalCostModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalCostModifiers = kv.NewSimpleMap[[]modifier.Modifier[CostContext]]()
		list := []modifier.Modifier[CostContext]{handler}
		m.addLocalCostModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddGlobalDirectAttackModifier(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalDirectAttackModifiers != nil {
		m.addGlobalDirectAttackModifiers = append(m.addGlobalDirectAttackModifiers, handler)
	} else {
		m.addGlobalDirectAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalFinalAttackModifier(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalFinalAttackModifiers != nil {
		m.addGlobalFinalAttackModifiers = append(m.addGlobalFinalAttackModifiers, handler)
	} else {
		m.addGlobalFinalAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalDefenceModifier(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalDefenceModifiers != nil {
		m.addGlobalDefenceModifiers = append(m.addGlobalDefenceModifiers, handler)
	} else {
		m.addGlobalDefenceModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalChargeModifier(handler modifier.Modifier[ChargeContext]) {
	if m.addGlobalChargeModifiers != nil {
		m.addGlobalChargeModifiers = append(m.addGlobalChargeModifiers, handler)
	} else {
		m.addGlobalChargeModifiers = []modifier.Modifier[ChargeContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalHealModifier(handler modifier.Modifier[HealContext]) {
	if m.addGlobalHealModifiers != nil {
		m.addGlobalHealModifiers = append(m.addGlobalHealModifiers, handler)
	} else {
		m.addGlobalHealModifiers = []modifier.Modifier[HealContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalCostModifier(handler modifier.Modifier[CostContext]) {
	if m.addGlobalCostModifiers != nil {
		m.addGlobalCostModifiers = append(m.addGlobalCostModifiers, handler)
	} else {
		m.addGlobalCostModifiers = []modifier.Modifier[CostContext]{handler}
	}
}

func (m *ModifierContext) RemoveLocalDirectAttackModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.removeLocalDirectAttackModifiers != null {
		modifiers := append(m.removeLocalDirectAttackModifiers.Get(targetCharacter), handler)
		m.removeLocalDirectAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalDirectAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalDirectAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalFinalAttackModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.removeLocalFinalAttackModifiers != null {
		modifiers := append(m.removeLocalFinalAttackModifiers.Get(targetCharacter), handler)
		m.removeLocalFinalAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalFinalAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalFinalAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalDefenceModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	var null kv.Map[uint, []modifier.Modifier[DamageContext]] = nil
	if m.removeLocalDefenceModifiers != null {
		modifiers := append(m.removeLocalDefenceModifiers.Get(targetCharacter), handler)
		m.removeLocalDefenceModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalDefenceModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalDefenceModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalChargeModifier(targetCharacter uint, handler modifier.Modifier[ChargeContext]) {
	var null kv.Map[uint, []modifier.Modifier[ChargeContext]] = nil
	if m.removeLocalChargeModifiers != null {
		modifiers := append(m.removeLocalChargeModifiers.Get(targetCharacter), handler)
		m.removeLocalChargeModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalChargeModifiers = kv.NewSimpleMap[[]modifier.Modifier[ChargeContext]]()
		list := []modifier.Modifier[ChargeContext]{handler}
		m.removeLocalChargeModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalHealModifier(targetCharacter uint, handler modifier.Modifier[HealContext]) {
	var null kv.Map[uint, []modifier.Modifier[HealContext]] = nil
	if m.removeLocalHealModifiers != null {
		modifiers := append(m.removeLocalHealModifiers.Get(targetCharacter), handler)
		m.removeLocalHealModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalHealModifiers = kv.NewSimpleMap[[]modifier.Modifier[HealContext]]()
		list := []modifier.Modifier[HealContext]{handler}
		m.removeLocalHealModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalCostModifier(targetCharacter uint, handler modifier.Modifier[CostContext]) {
	var null kv.Map[uint, []modifier.Modifier[CostContext]] = nil
	if m.removeLocalCostModifiers != null {
		modifiers := append(m.removeLocalCostModifiers.Get(targetCharacter), handler)
		m.removeLocalCostModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalCostModifiers = kv.NewSimpleMap[[]modifier.Modifier[CostContext]]()
		list := []modifier.Modifier[CostContext]{handler}
		m.removeLocalCostModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveGlobalDirectAttackModifier(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalDirectAttackModifiers != nil {
		m.removeGlobalDirectAttackModifiers = append(m.removeGlobalDirectAttackModifiers, handler)
	} else {
		m.removeGlobalDirectAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalFinalAttackModifier(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalFinalAttackModifiers != nil {
		m.removeGlobalFinalAttackModifiers = append(m.removeGlobalFinalAttackModifiers, handler)
	} else {
		m.removeGlobalFinalAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalDefenceModifier(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalDefenceModifiers != nil {
		m.removeGlobalDefenceModifiers = append(m.removeGlobalDefenceModifiers, handler)
	} else {
		m.removeGlobalDefenceModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalChargeModifier(handler modifier.Modifier[ChargeContext]) {
	if m.removeGlobalChargeModifiers != nil {
		m.removeGlobalChargeModifiers = append(m.removeGlobalChargeModifiers, handler)
	} else {
		m.removeGlobalChargeModifiers = []modifier.Modifier[ChargeContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalHealModifier(handler modifier.Modifier[HealContext]) {
	if m.removeGlobalHealModifiers != nil {
		m.removeGlobalHealModifiers = append(m.removeGlobalHealModifiers, handler)
	} else {
		m.removeGlobalHealModifiers = []modifier.Modifier[HealContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalCostModifier(handler modifier.Modifier[CostContext]) {
	if m.removeGlobalCostModifiers != nil {
		m.removeGlobalCostModifiers = append(m.removeGlobalCostModifiers, handler)
	} else {
		m.removeGlobalCostModifiers = []modifier.Modifier[CostContext]{handler}
	}
}

func (m ModifierContext) AddLocalDirectAttackModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.addLocalDirectAttackModifiers
}

func (m ModifierContext) AddLocalFinalAttackModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.addLocalFinalAttackModifiers
}

func (m ModifierContext) AddLocalDefenceModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.addLocalDefenceModifiers
}

func (m ModifierContext) AddLocalChargeModifiers() kv.Map[uint, []modifier.Modifier[ChargeContext]] {
	return m.addLocalChargeModifiers
}

func (m ModifierContext) AddLocalHealModifiers() kv.Map[uint, []modifier.Modifier[HealContext]] {
	return m.addLocalHealModifiers
}

func (m ModifierContext) AddLocalCostModifiers() kv.Map[uint, []modifier.Modifier[CostContext]] {
	return m.addLocalCostModifiers
}

func (m ModifierContext) AddGlobalDirectAttackModifiers() []modifier.Modifier[DamageContext] {
	return m.addGlobalDirectAttackModifiers
}

func (m ModifierContext) AddGlobalFinalAttackModifiers() []modifier.Modifier[DamageContext] {
	return m.addGlobalFinalAttackModifiers
}

func (m ModifierContext) AddGlobalDefenceModifiers() []modifier.Modifier[DamageContext] {
	return m.addGlobalDefenceModifiers
}

func (m ModifierContext) AddGlobalChargeModifiers() []modifier.Modifier[ChargeContext] {
	return m.addGlobalChargeModifiers
}

func (m ModifierContext) AddGlobalHealModifiers() []modifier.Modifier[HealContext] {
	return m.addGlobalHealModifiers
}

func (m ModifierContext) AddGlobalCostModifiers() []modifier.Modifier[CostContext] {
	return m.addGlobalCostModifiers
}

func (m ModifierContext) RemoveLocalDirectAttackModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.removeLocalDirectAttackModifiers
}

func (m ModifierContext) RemoveLocalFinalAttackModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.removeLocalFinalAttackModifiers
}

func (m ModifierContext) RemoveLocalDefenceModifiers() kv.Map[uint, []modifier.Modifier[DamageContext]] {
	return m.removeLocalDefenceModifiers
}

func (m ModifierContext) RemoveLocalChargeModifiers() kv.Map[uint, []modifier.Modifier[ChargeContext]] {
	return m.removeLocalChargeModifiers
}

func (m ModifierContext) RemoveLocalHealModifiers() kv.Map[uint, []modifier.Modifier[HealContext]] {
	return m.removeLocalHealModifiers
}

func (m ModifierContext) RemoveLocalCostModifiers() kv.Map[uint, []modifier.Modifier[CostContext]] {
	return m.removeLocalCostModifiers
}

func (m ModifierContext) RemoveGlobalDirectAttackModifiers() []modifier.Modifier[DamageContext] {
	return m.removeGlobalDirectAttackModifiers
}

func (m ModifierContext) RemoveGlobalFinalAttackModifiers() []modifier.Modifier[DamageContext] {
	return m.removeGlobalFinalAttackModifiers
}

func (m ModifierContext) RemoveGlobalDefenceModifiers() []modifier.Modifier[DamageContext] {
	return m.removeGlobalDefenceModifiers
}

func (m ModifierContext) RemoveGlobalChargeModifiers() []modifier.Modifier[ChargeContext] {
	return m.removeGlobalChargeModifiers
}

func (m ModifierContext) RemoveGlobalHealModifiers() []modifier.Modifier[HealContext] {
	return m.removeGlobalHealModifiers
}

func (m ModifierContext) RemoveGlobalCostModifiers() []modifier.Modifier[CostContext] {
	return m.removeGlobalCostModifiers
}
