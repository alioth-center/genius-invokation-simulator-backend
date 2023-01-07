package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

// Card 被持久化模块托管的Card工厂
type Card struct {
	Card entity.Card
}

// Skill 被持久化模块托管的Skill工厂
type Skill struct {
	Skill entity.Skill
}

// RuleSet 被持久化模块托管的RuleSet工厂
type RuleSet struct {
	Rule entity.RuleSet
}

// Character 被持久化模块托管的CharacterInfo工厂
type Character struct {
	ID          uint
	Affiliation enum.Affiliation
	Vision      enum.ElementType
	Weapon      enum.WeaponType
	MaxHP       uint
	MaxMP       uint
	Skills      []uint
}

// Player 被持久化模块托管的PlayerInfo工厂
type Player struct {
	UID       uint
	NickName  string
	CardDecks []uint
	Password  string
}

// CardDeck 被持久化模块托管的CardDeck工厂
type CardDeck struct {
	ID         uint
	OwnerUID   uint
	Cards      []uint
	Characters []uint
}
