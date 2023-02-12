package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
)

type cacheableImpl struct {
	id uint64
}

func (c cacheableImpl) ID() uint64 {
	return c.id
}

func newCacheableImpl(id uint64) cacheableImpl {
	return cacheableImpl{id: id}
}

type CardPersistenceAdapter struct{}

func (c CardPersistenceAdapter) Convert(source Card) (success bool, result persistence.Card) {
	return true, persistence.Card{
		Cacheable: newCacheableImpl(source.TypeID()),
		Card:      source,
	}
}

func NewCardPersistenceAdapter() adapter.Adapter[Card, persistence.Card] {
	return CardPersistenceAdapter{}
}

type SkillPersistenceAdapter struct{}

func (s SkillPersistenceAdapter) Convert(source Skill) (success bool, result persistence.Skill) {
	return true, persistence.Skill{
		Cacheable: newCacheableImpl(source.TypeID()),
		Skill:     source,
	}
}

func NewSkillPersistenceAdapter() adapter.Adapter[Skill, persistence.Skill] {
	return SkillPersistenceAdapter{}
}

type EventPersistenceAdapter struct{}

func (e EventPersistenceAdapter) Convert(source Event) (success bool, result persistence.Event) {
	return true, persistence.Event{
		Cacheable: newCacheableImpl(source.TypeID()),
		Event:     source,
	}
}

func NewEventPersistenceAdapter() adapter.Adapter[Event, persistence.Event] {
	return EventPersistenceAdapter{}
}

type RuleSetPersistenceAdapter struct{}

func (r RuleSetPersistenceAdapter) Convert(source RuleSet) (success bool, result persistence.RuleSet) {
	panic("not implemented")
	return true, persistence.RuleSet{
		Cacheable: newCacheableImpl(0),
		Rule:      source,
	}
}

func NewRuleSetAdapter() adapter.Adapter[RuleSet, persistence.RuleSet] {
	return RuleSetPersistenceAdapter{}
}
