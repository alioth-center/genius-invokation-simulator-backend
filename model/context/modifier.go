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
	if m.addLocalDirectAttackModifiers != nil {
		modifiers := append(m.addLocalDirectAttackModifiers.Get(targetCharacter), handler)
		m.addLocalDirectAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalDirectAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalDirectAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalFinalAttackModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	if m.addLocalFinalAttackModifiers != nil {
		modifiers := append(m.addLocalFinalAttackModifiers.Get(targetCharacter), handler)
		m.addLocalFinalAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalFinalAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalFinalAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalDefenceModifier(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	if m.addLocalDefenceModifiers != nil {
		modifiers := append(m.addLocalDefenceModifiers.Get(targetCharacter), handler)
		m.addLocalDefenceModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalDefenceModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.addLocalDefenceModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalChargeModifier(targetCharacter uint, handler modifier.Modifier[ChargeContext]) {
	if m.addLocalChargeModifiers != nil {
		modifiers := append(m.addLocalChargeModifiers.Get(targetCharacter), handler)
		m.addLocalChargeModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalChargeModifiers = kv.NewSimpleMap[[]modifier.Modifier[ChargeContext]]()
		list := []modifier.Modifier[ChargeContext]{handler}
		m.addLocalChargeModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalHealModifier(targetCharacter uint, handler modifier.Modifier[HealContext]) {
	if m.addLocalHealModifiers != nil {
		modifiers := append(m.addLocalHealModifiers.Get(targetCharacter), handler)
		m.addLocalHealModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalHealModifiers = kv.NewSimpleMap[[]modifier.Modifier[HealContext]]()
		list := []modifier.Modifier[HealContext]{handler}
		m.addLocalHealModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddLocalCostModifier(targetCharacter uint, handler modifier.Modifier[CostContext]) {
	if m.addLocalCostModifiers != nil {
		modifiers := append(m.addLocalCostModifiers.Get(targetCharacter), handler)
		m.addLocalCostModifiers.Set(targetCharacter, modifiers)
	} else {
		m.addLocalCostModifiers = kv.NewSimpleMap[[]modifier.Modifier[CostContext]]()
		list := []modifier.Modifier[CostContext]{handler}
		m.addLocalCostModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) AddGlobalDirectAttackModifiers(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalDirectAttackModifiers != nil {
		m.addGlobalDirectAttackModifiers = append(m.addGlobalDirectAttackModifiers, handler)
	} else {
		m.addGlobalDirectAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalFinalAttackModifiers(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalFinalAttackModifiers != nil {
		m.addGlobalFinalAttackModifiers = append(m.addGlobalFinalAttackModifiers, handler)
	} else {
		m.addGlobalFinalAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalDefenceModifiers(handler modifier.Modifier[DamageContext]) {
	if m.addGlobalDefenceModifiers != nil {
		m.addGlobalDefenceModifiers = append(m.addGlobalDefenceModifiers, handler)
	} else {
		m.addGlobalDefenceModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalChargeModifiers(handler modifier.Modifier[ChargeContext]) {
	if m.addGlobalChargeModifiers != nil {
		m.addGlobalChargeModifiers = append(m.addGlobalChargeModifiers, handler)
	} else {
		m.addGlobalChargeModifiers = []modifier.Modifier[ChargeContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalHealModifiers(handler modifier.Modifier[HealContext]) {
	if m.addGlobalHealModifiers != nil {
		m.addGlobalHealModifiers = append(m.addGlobalHealModifiers, handler)
	} else {
		m.addGlobalHealModifiers = []modifier.Modifier[HealContext]{handler}
	}
}

func (m *ModifierContext) AddGlobalCostModifiers(handler modifier.Modifier[CostContext]) {
	if m.addGlobalCostModifiers != nil {
		m.addGlobalCostModifiers = append(m.addGlobalCostModifiers, handler)
	} else {
		m.addGlobalCostModifiers = []modifier.Modifier[CostContext]{handler}
	}
}

func (m *ModifierContext) RemoveLocalDirectAttackModifiers(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	if m.removeLocalDirectAttackModifiers != nil {
		modifiers := append(m.removeLocalDirectAttackModifiers.Get(targetCharacter), handler)
		m.removeLocalDirectAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalDirectAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalDirectAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalFinalAttackModifiers(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	if m.removeLocalFinalAttackModifiers != nil {
		modifiers := append(m.removeLocalFinalAttackModifiers.Get(targetCharacter), handler)
		m.removeLocalFinalAttackModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalFinalAttackModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalFinalAttackModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalDefenceModifiers(targetCharacter uint, handler modifier.Modifier[DamageContext]) {
	if m.removeLocalDefenceModifiers != nil {
		modifiers := append(m.removeLocalDefenceModifiers.Get(targetCharacter), handler)
		m.removeLocalDefenceModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalDefenceModifiers = kv.NewSimpleMap[[]modifier.Modifier[DamageContext]]()
		list := []modifier.Modifier[DamageContext]{handler}
		m.removeLocalDefenceModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalChargeModifiers(targetCharacter uint, handler modifier.Modifier[ChargeContext]) {
	if m.removeLocalChargeModifiers != nil {
		modifiers := append(m.removeLocalChargeModifiers.Get(targetCharacter), handler)
		m.removeLocalChargeModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalChargeModifiers = kv.NewSimpleMap[[]modifier.Modifier[ChargeContext]]()
		list := []modifier.Modifier[ChargeContext]{handler}
		m.removeLocalChargeModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalHealModifiers(targetCharacter uint, handler modifier.Modifier[HealContext]) {
	if m.removeLocalHealModifiers != nil {
		modifiers := append(m.removeLocalHealModifiers.Get(targetCharacter), handler)
		m.removeLocalHealModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalHealModifiers = kv.NewSimpleMap[[]modifier.Modifier[HealContext]]()
		list := []modifier.Modifier[HealContext]{handler}
		m.removeLocalHealModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveLocalCostModifiers(targetCharacter uint, handler modifier.Modifier[CostContext]) {
	if m.removeLocalCostModifiers != nil {
		modifiers := append(m.removeLocalCostModifiers.Get(targetCharacter), handler)
		m.removeLocalCostModifiers.Set(targetCharacter, modifiers)
	} else {
		m.removeLocalCostModifiers = kv.NewSimpleMap[[]modifier.Modifier[CostContext]]()
		list := []modifier.Modifier[CostContext]{handler}
		m.removeLocalCostModifiers.Set(targetCharacter, list)
	}
}

func (m *ModifierContext) RemoveGlobalDirectAttackModifiers(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalDirectAttackModifiers != nil {
		m.removeGlobalDirectAttackModifiers = append(m.removeGlobalDirectAttackModifiers, handler)
	} else {
		m.removeGlobalDirectAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalFinalAttackModifiers(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalFinalAttackModifiers != nil {
		m.removeGlobalFinalAttackModifiers = append(m.removeGlobalFinalAttackModifiers, handler)
	} else {
		m.removeGlobalFinalAttackModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalDefenceModifiers(handler modifier.Modifier[DamageContext]) {
	if m.removeGlobalDefenceModifiers != nil {
		m.removeGlobalDefenceModifiers = append(m.removeGlobalDefenceModifiers, handler)
	} else {
		m.removeGlobalDefenceModifiers = []modifier.Modifier[DamageContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalChargeModifiers(handler modifier.Modifier[ChargeContext]) {
	if m.removeGlobalChargeModifiers != nil {
		m.removeGlobalChargeModifiers = append(m.removeGlobalChargeModifiers, handler)
	} else {
		m.removeGlobalChargeModifiers = []modifier.Modifier[ChargeContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalHealModifiers(handler modifier.Modifier[HealContext]) {
	if m.removeGlobalHealModifiers != nil {
		m.removeGlobalHealModifiers = append(m.removeGlobalHealModifiers, handler)
	} else {
		m.removeGlobalHealModifiers = []modifier.Modifier[HealContext]{handler}
	}
}

func (m *ModifierContext) RemoveGlobalCostModifiers(handler modifier.Modifier[CostContext]) {
	if m.removeGlobalCostModifiers != nil {
		m.removeGlobalCostModifiers = append(m.removeGlobalCostModifiers, handler)
	} else {
		m.removeGlobalCostModifiers = []modifier.Modifier[CostContext]{handler}
	}
}
