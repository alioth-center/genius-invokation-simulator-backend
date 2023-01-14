package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	adapter "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	converter "github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
)

type cardAdapterLayer struct {
	implement.BaseEntityImplement
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

func (c CardAdapter) Convert(source adapter.Card) (success bool, result model.Card) {
	adapterLayer := &cardAdapterLayer{
		BaseEntityImplement: implement.BaseEntityImplement{},
		cardType:            source.CardType(),
		cost:                source.Cost(),
	}

	return true, adapterLayer
}

func NewCardAdapter() converter.Adapter[adapter.Card, model.Card] {
	return CardAdapter{}
}

type eventCardAdapterLayer struct {
	cardAdapterLayer
	event adapter.Event
}

func (e *eventCardAdapterLayer) Event() (event model.Event) {
	_, event = eventAdapter.Convert(e.event)
	return event
}

type EventCardAdapter struct{}

func (e EventCardAdapter) Convert(source adapter.EventCard) (success bool, result model.EventCard) {
	adapterLayer := &eventCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImplement: implement.BaseEntityImplement{},
			cardType:            source.CardType(),
			cost:                source.Cost(),
		},
		event: source.Event(),
	}

	return true, adapterLayer
}

func NewEventCardAdapter() converter.Adapter[adapter.EventCard, model.EventCard] {
	return EventCardAdapter{}
}

type supportCardAdapterLayer struct {
	cardAdapterLayer
	event adapter.Event
}

func (s *supportCardAdapterLayer) Support() (event model.Event) {
	_, event = eventAdapter.Convert(s.event)
	return event
}

type SupportCardAdapter struct{}

func (s SupportCardAdapter) Convert(source adapter.SupportCard) (success bool, result model.SupportCard) {
	adapterLayer := &supportCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImplement: implement.BaseEntityImplement{},
			cardType:            source.CardType(),
			cost:                source.Cost(),
		},
		event: source.Support(),
	}

	return true, adapterLayer
}

func NewSupportCardAdapter() converter.Adapter[adapter.SupportCard, model.SupportCard] {
	return SupportCardAdapter{}
}

type equipmentCardAdapterLayer struct {
	cardAdapterLayer
	event         adapter.Event
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

func (e EquipmentCardAdapter) Convert(source adapter.EquipmentCard) (success bool, result model.EquipmentCard) {
	adapterLayer := &equipmentCardAdapterLayer{
		cardAdapterLayer: cardAdapterLayer{
			BaseEntityImplement: implement.BaseEntityImplement{},
			cardType:            source.CardType(),
			cost:                source.Cost(),
		},
		event:         source.Modify(),
		equipmentType: source.EquipmentType(),
	}

	return true, adapterLayer
}

func NewEquipmentCardAdapter() converter.Adapter[adapter.EquipmentCard, model.EquipmentCard] {
	return EquipmentCardAdapter{}
}
