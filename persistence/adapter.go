package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
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

func (c CardPersistenceAdapter) Convert(source model.Card) (success bool, result Card) {
	return true, Card{
		Cacheable: newCacheableImpl(source.TypeID()),
		Card:      source,
	}
}

func NewCardPersistenceAdapter() adapter.Adapter[model.Card, Card] {
	return CardPersistenceAdapter{}
}

type SkillPersistenceAdapter struct{}

func (s SkillPersistenceAdapter) Convert(source model.Skill) (success bool, result Skill) {
	return true, Skill{
		Cacheable: newCacheableImpl(source.TypeID()),
		Skill:     source,
	}
}

func NewSkillPersistenceAdapter() adapter.Adapter[model.Skill, Skill] {
	return SkillPersistenceAdapter{}
}

type EventPersistenceAdapter struct{}

func (e EventPersistenceAdapter) Convert(source model.Event) (success bool, result Event) {
	return true, Event{
		Cacheable: newCacheableImpl(source.TypeID()),
		Event:     source,
	}
}

func NewEventPersistenceAdapter() adapter.Adapter[model.Event, Event] {
	return EventPersistenceAdapter{}
}

type RuleSetPersistenceAdapter struct{}

func (r RuleSetPersistenceAdapter) Convert(source model.RuleSet) (success bool, result RuleSet) {
	return true, RuleSet{
		Cacheable: newCacheableImpl(source.ID),
		Rule:      source,
	}
}

func NewRuleSetAdapter() adapter.Adapter[model.RuleSet, RuleSet] {
	return RuleSetPersistenceAdapter{}
}
