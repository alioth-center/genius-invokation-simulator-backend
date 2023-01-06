package message

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

// SyncMessageInterface 同步信息类型的约束接口
type SyncMessageInterface interface {
	PlayerMessage | ViewerMessage | GuestMessage
}

// DictionaryPair 追加字典，将同步信息中给的实体ID与它们的类型ID相联系
type DictionaryPair struct {
	TypeName string `json:"type_name" yaml:"type_name" xml:"type_name"` // TypeName 类型的名称
	TypeID   uint   `json:"type_id" yaml:"type_id" xml:"type_id"`       // TypeID 类型的ID
	EntityID uint   `json:"entity_id" yaml:"entity_id" xml:"entity_id"` // EntityID 实体的ID
}

// Modifier 修正BUFF
type Modifier struct {
	ID   uint              `json:"id" yaml:"id" xml:"id"`       // ID 修正BUFF的实体ID
	Type enum.ModifierType `json:"type" yaml:"type" xml:"type"` // Type 修正BUFF的类型
}

// Equipment 装备信息
type Equipment struct {
	ID   uint               `json:"id" yaml:"id" xml:"id"`       // ID 装备的实体ID
	Type enum.EquipmentType `json:"type" yaml:"type" xml:"type"` // Type 装备的类型
}

// CooperativeSkill 协同攻击
type CooperativeSkill struct {
	ID      uint             `json:"id" yaml:"id" xml:"id"`                // ID 协同攻击的实体ID
	Trigger enum.TriggerType `json:"trigger" yaml:"trigger" xml:"trigger"` // Trigger 协同攻击的触发条件
}

// Event 事件信息
type Event struct {
	ID      uint             `json:"id" yaml:"id" xml:"id"`                // ID 事件的实体ID
	Trigger enum.TriggerType `json:"trigger" yaml:"trigger" xml:"trigger"` // Trigger 事件的触发条件
}

// Character 角色信息
type Character struct {
	ID         uint                 `json:"id" yaml:"id" xml:"id"`                         // ID 角色的实体ID
	MP         uint                 `json:"mp" yaml:"mp" xml:"mp"`                         // MP 角色的当前充能
	HP         uint                 `json:"hp" yaml:"hp" xml:"hp"`                         // HP 角色的当前生命
	Equipments []Equipment          `json:"equipments" yaml:"equipments" xml:"equipments"` // Equipment 角色当前的装备
	Modifiers  []Modifier           `json:"modifiers" yaml:"modifiers" xml:"modifiers"`    // Modifiers 角色当前的修正BUFF
	Status     enum.CharacterStatus `json:"status" yaml:"status" xml:"status"`             // Status 角色当前的状态
}

// Base 基础玩家信息
type Base struct {
	UID          uint               `json:"uid" yaml:"uid" xml:"uid"`                            // UID 玩家的UID
	Characters   []Character        `json:"characters" yaml:"characters" xml:"characters"`       // Characters 玩家的持有角色
	CampEffect   []Modifier         `json:"camp_effect" yaml:"camp_effect" xml:"camp_effect"`    // CampEffect 玩家的阵营效果
	Cooperatives []CooperativeSkill `json:"cooperatives" yaml:"cooperatives" xml:"cooperatives"` // Cooperatives 玩家可进行的协同攻击
	Events       []Event            `json:"events" yaml:"events" xml:"events"`                   // Events 玩家身上的事件
	RemainCards  uint               `json:"remain_cards" yaml:"remain_cards" xml:"remain_cards"` // RemainCards 玩家牌堆剩余的数量
	Status       enum.PlayerStatus  `json:"status" yaml:"status" xml:"status"`                   // Status 玩家的状态信息
}

// Self 接收玩家自己的信息
type Self struct {
	Base
	Cost  map[enum.ElementType]uint `json:"cost" yaml:"cost" xml:"cost"`    // Cost 玩家持有的元素骰子
	Cards []uint                    `json:"cards" yaml:"cards" xml:"cards"` // Cards 玩家持有的卡牌
}

// Other 接收玩家所见的其他玩家信息
type Other struct {
	Base
	Cost  uint `json:"cost"  yaml:"cost" xml:"cost"`    // Cost 玩家持有的元素骰子数量
	Cards uint `json:"cards"  yaml:"cards" xml:"cards"` // Cards 玩家持有的卡牌数量
}

// SyncMessage 玩家接收到的同步消息
type SyncMessage struct {
	Target  uint        `json:"target"   yaml:"target" xml:"target"`  // Target 接收同步消息的玩家
	Message interface{} `json:"message" yaml:"message" xml:"message"` // Message 同步消息
}

// PlayerMessage 参与玩家接收到的同步信息
type PlayerMessage struct {
	Self   Self             `json:"self" yaml:"self" xml:"self"`       // Self 自己的信息
	Others []Other          `json:"others" yaml:"others" xml:"others"` // Others 其他人的信息
	Append []DictionaryPair `json:"append" yaml:"append" xml:"append"` // Append 追加的字典
}

// ViewerMessage 观战玩家接收到的同步信息
type ViewerMessage struct {
	Players []Self           `json:"players" yaml:"players" xml:"players"` // Players 在场玩家的信息
	Append  []DictionaryPair `json:"append" yaml:"append" xml:"append"`    // Append 追加的字典
}

// GuestMessage 游客接收道德同步信息
type GuestMessage struct {
	Players []Other          `json:"players" yaml:"players" xml:"players"` // Players 在场玩家的信息
	Append  []DictionaryPair `json:"append" yaml:"append" xml:"append"`    // Append 追加的字典
}

// NewSyncMessage 创建一个指定接收者的同步信息
func NewSyncMessage[message SyncMessageInterface](target uint, msg message) SyncMessage {
	return SyncMessage{
		Target:  target,
		Message: msg,
	}
}
