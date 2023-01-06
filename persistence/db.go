package persistence

import "github.com/sunist-c/genius-invokation-simulator-backend/entity"

var (
	RuleSetPersistence   Persistence[entity.RuleSet]
	CardPersistence      Persistence[entity.Card]
	CharacterPersistence Persistence[entity.Character]
	SkillPersistence     Persistence[entity.Skill]
)

type Persistence[T any] struct {
}

func (p Persistence[T]) QueryByID(id uint) (has bool, result T) {
	return false, result
}
