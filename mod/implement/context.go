package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type ContextImpl struct {
	activeCharacter      model.Character
	backgroundCharacters []model.Character
	self                 model.Player
	players              []model.Player
}

func (impl *ContextImpl) ActiveCharacter() model.Character {
	return impl.activeCharacter
}

func (impl *ContextImpl) BackgroundCharacters() []model.Character {
	return impl.backgroundCharacters
}

func (impl *ContextImpl) Self() model.Player {
	return impl.self
}

func (impl *ContextImpl) Players() []model.Player {
	return impl.players
}

func NewEmptyContext() definition.Context {
	return &ContextImpl{
		activeCharacter:      nil,
		backgroundCharacters: []model.Character{},
		self:                 nil,
		players:              []model.Player{},
	}
}
