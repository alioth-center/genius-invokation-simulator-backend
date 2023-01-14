package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	definition "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
)

type cardAdapterLayer struct {
	implement.BaseEntityImpl
	cardType enum.CardType
	cost     map[enum.ElementType]uint
}

func (c *cardAdapterLayer) Type() enum.CardType {
	return c.cardType
}

func (c *cardAdapterLayer) Cost() map[enum.ElementType]uint {
	return c.cost
}

type CardAdapter struct{}

func (c CardAdapter) Convert(source definition.Card) (success bool, result model.Card) {
	adapterLayer := &cardAdapterLayer{
		BaseEntityImpl: implement.BaseEntityImpl{},
		cardType:       source.CardType(),
		cost:           source.Cost(),
	}

	return true, adapterLayer
}

func NewCardAdapter() adapter.Adapter[definition.Card, model.Card] {
	return CardAdapter{}
}

type eventCardAdapterLayer struct {
	cardAdapterLayer
	event definition.Event
}

func (e *eventCardAdapterLayer) Event() (event model.Event) {
	_, event = eventAdapter.Convert(e.event)
	return event
}

type EventCardAdapter struct{}

func (e EventCardAdapter) Convert(source definition.EventCard) (success bool, result model.EventCard) {
	adapterLayer := &eventCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImpl: implement.BaseEntityImpl{},
			cardType:       source.CardType(),
			cost:           source.Cost(),
		},
		event: source.Event(),
	}

	return true, adapterLayer
}

func NewEventCardAdapter() adapter.Adapter[definition.EventCard, model.EventCard] {
	return EventCardAdapter{}
}

type supportCardAdapterLayer struct {
	cardAdapterLayer
	event definition.Event
}

func (s *supportCardAdapterLayer) Support() (event model.Event) {
	_, event = eventAdapter.Convert(s.event)
	return event
}

type SupportCardAdapter struct{}

func (s SupportCardAdapter) Convert(source definition.SupportCard) (success bool, result model.SupportCard) {
	adapterLayer := &supportCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImpl: implement.BaseEntityImpl{},
			cardType:       source.CardType(),
			cost:           source.Cost(),
		},
		event: source.Support(),
	}

	return true, adapterLayer
}

func NewSupportCardAdapter() adapter.Adapter[definition.SupportCard, model.SupportCard] {
	return SupportCardAdapter{}
}

type equipmentCardAdapterLayer struct {
	cardAdapterLayer
	event         definition.Event
	equipmentType enum.EquipmentType
}

func (e *equipmentCardAdapterLayer) EquipmentType() enum.EquipmentType {
	return e.equipmentType
}

func (e *equipmentCardAdapterLayer) Modify() (event model.Event) {
	_, event = eventAdapter.Convert(e.event)
	return event
}

type EquipmentCardAdapter struct{}

func (e EquipmentCardAdapter) Convert(source definition.EquipmentCard) (success bool, result model.EquipmentCard) {
	adapterLayer := &equipmentCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImpl: implement.BaseEntityImpl{},
			cardType:       source.CardType(),
			cost:           source.Cost(),
		},
		event:         source.Modify(),
		equipmentType: source.EquipmentType(),
	}

	return true, adapterLayer
}

func NewEquipmentCardAdapter() adapter.Adapter[definition.EquipmentCard, model.EquipmentCard] {
	return EquipmentCardAdapter{}
}

type weaponCardAdapterLayer struct {
	equipmentCardAdapterLayer
	weaponType enum.WeaponType
}

func (w *weaponCardAdapterLayer) WeaponType() enum.WeaponType {
	return w.weaponType
}

type WeaponCardAdapter struct{}

func (w WeaponCardAdapter) Convert(source definition.WeaponCard) (success bool, result model.WeaponCard) {
	adapterLayer := &weaponCardAdapterLayer{
		equipmentCardAdapterLayer: equipmentCardAdapterLayer{
			cardAdapterLayer: cardAdapterLayer{
				BaseEntityImpl: implement.BaseEntityImpl{},
				cardType:       source.CardType(),
				cost:           source.Cost(),
			},
			event:         source.Modify(),
			equipmentType: source.EquipmentType(),
		},
		weaponType: source.WeaponType(),
	}

	return true, adapterLayer
}

func NewWeaponCardAdapter() adapter.Adapter[definition.WeaponCard, model.WeaponCard] {
	return WeaponCardAdapter{}
}
