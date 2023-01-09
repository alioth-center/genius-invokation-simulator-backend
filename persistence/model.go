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
	UID       uint   `xorm:"pk autoincr notnull unique index"` // UID Player的UID，主键
	NickName  string `xorm:"notnull varchar(64)"`              // NickName Player的昵称
	CardDecks []uint `xorm:"notnull json"`                     // CardDecks Player保存的卡组
	Password  string `xorm:"notnull varchar(64)"`              // Password Player的密码Hash
}

// CardDeck 被持久化模块托管的CardDeck工厂
type CardDeck struct {
	ID               uint     `xorm:"pk autoincr notnull unique index"` // ID CardDeck的记录ID，主键
	OwnerUID         uint     `xorm:"notnull index"`                    // OwnerUID CardDeck的持有者
	RequiredPackages []string `xorm:"notnull json"`                     // RequiredPackages CardDeck需要的包
	Cards            []uint   `xorm:"notnull json"`                     // Cards CardDeck包含的卡组
	Characters       []uint   `xorm:"notnull json"`                     // Characters CardDeck包含的角色
}
