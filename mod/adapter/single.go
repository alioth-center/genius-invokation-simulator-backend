package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
)

var (
	eventAdapter = NewEventAdapter()

	cardAdapter          = NewCardAdapter()
	eventCardAdapter     = NewEventCardAdapter()
	supportCardAdapter   = NewSupportCardAdapter()
	equipmentCardAdapter = NewEquipmentCardAdapter()
	weaponCardAdapter    = NewWeaponCardAdapter()
	ruleSetAdapter       = NewRuleSetAdapter()

	attackModifierAdapter = NewAttackModifierAdapter()
	costModifierAdapter   = NewCostModifierAdapter()
	healModifierAdapter   = NewHealModifierAdapter()
	chargeModifierAdapter = NewChargeModifierAdapter()
)

func GetEventAdapter() adapter.Adapter[definition.Event, model.Event] {
	return eventAdapter
}

func GetCardAdapter() adapter.Adapter[definition.Card, model.Card] {
	return cardAdapter
}

func GetEventCardAdapter() adapter.Adapter[definition.EventCard, model.EventCard] {
	return eventCardAdapter
}

func GetSupportCardAdapter() adapter.Adapter[definition.SupportCard, model.SupportCard] {
	return supportCardAdapter
}

func GetEquipmentCardAdapter() adapter.Adapter[definition.EquipmentCard, model.EquipmentCard] {
	return equipmentCardAdapter
}

func GetWeaponCardAdapter() adapter.Adapter[definition.WeaponCard, model.WeaponCard] {
	return weaponCardAdapter
}

func GetRuleSetAdapter() adapter.Adapter[definition.Rule, model.RuleSet] {
	return ruleSetAdapter
}
