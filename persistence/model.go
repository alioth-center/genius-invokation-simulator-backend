package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/localization"
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
	UID       uint   `xorm:"pk autoincr notnull unique index"`
	NickName  string `xorm:"notnull varchar(64)"`
	CardDecks []uint `xorm:"notnull json"`
	Password  string `xorm:"notnull varchar(64)"`
}

// CardDeck 被持久化模块托管的CardDeck工厂
type CardDeck struct {
	ID         uint   `xorm:"pk autoincr notnull unique index"`
	OwnerUID   uint   `xorm:"notnull index"`
	Cards      []uint `xorm:"notnull json"`
	Characters []uint `xorm:"notnull json"`
}

// Localization 被持久化模块托管的本地化信息工厂
type Localization struct {
	EntityID         uint
	EntityUID        string
	MultipleLanguage localization.LanguagePack
}
