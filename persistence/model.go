package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/websocket/message"
	"time"
)

// Card 被持久化模块托管的Card工厂的产品
type Card struct {
	Cacheable
	Card model.Card
}

// Skill 被持久化模块托管的Skill工厂的产品
type Skill struct {
	Cacheable
	Skill model.Skill
}

// RuleSet 被持久化模块托管的RuleSet工厂的产品
type RuleSet struct {
	Cacheable
	Rule model.RuleSet
}

// Summon 被持久化模块托管的Summon工厂的产品
type Summon struct {
	Cacheable
}

// Event 被持久化模块托管的Event工厂的产品
type Event struct {
	Cacheable
}

// Character 被持久化模块托管的CharacterInfo工厂的产品
type Character struct {
	Cacheable
	EntityID    uint
	Affiliation enum.Affiliation
	Vision      enum.ElementType
	Weapon      enum.WeaponType
	MaxHP       uint
	MaxMP       uint
	Skills      []uint
}

func (c Character) ID() uint { return c.EntityID }

// Player 被持久化模块托管的Player信息
type Player struct {
	UID      uint   `xorm:"pk autoincr notnull unique index"` // UID Player的UID，主键
	NickName string `xorm:"notnull varchar(64)"`              // NickName Player的昵称
	Password string `xorm:"notnull varchar(64)"`              // Password Player的密码Hash
}

// CardDeck 被持久化模块托管的CardDeck信息
type CardDeck struct {
	ID               uint     `xorm:"pk autoincr notnull unique index" json:"id"` // ID CardDeck的记录ID，主键
	OwnerUID         uint     `xorm:"notnull index" json:"owner_uid"`             // OwnerUID CardDeck的持有者
	RequiredPackages []string `xorm:"notnull json" json:"required_packages"`      // RequiredPackages CardDeck需要的包
	Cards            []uint   `xorm:"notnull json" json:"cards"`                  // Cards CardDeck包含的卡组
	Characters       []uint   `xorm:"notnull json" json:"characters"`             // Characters CardDeck包含的角色
}

// Token 被持久化模块托管的Token缓存
type Token struct {
	UID uint   // UID Token持有者的UID
	ID  string // ID Token的ID
}

// RoomInfo 被持久化模块托管的RoomInfo缓存
type RoomInfo struct {
	RoomID           uint                `json:"room_id"`
	CreatedAt        time.Time           `json:"created_at"`
	CreatorID        uint                `json:"creator_id"`
	Token            string              `json:"token"`
	Players          []uint              `json:"players"`
	Viewers          uint                `json:"viewers"`
	RequiredPackages []string            `json:"required_packages"`
	GameOptions      message.GameOptions `json:"game_options"`
}

// ModInfo 被持久化模块托管的ModInfo缓存
type ModInfo struct {
	PackageName string `json:"package_name"`
	Events      []uint `json:"events"`
	Characters  []uint `json:"characters"`
	Cards       []uint `json:"cards"`
	Skills      []uint `json:"skills"`
	Summons     []uint `json:"summons"`
}
