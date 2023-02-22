package definition

import "github.com/sunist-c/genius-invokation-simulator-backend/entity/model"

type Context interface {
	ActiveCharacter() model.Character
	BackgroundCharacters() []model.Character
	Self() model.Player
	Players() []model.Player
}
