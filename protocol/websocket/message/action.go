package message

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

// ActionMessageInterface 玩家操作信息的约束接口
type ActionMessageInterface interface {
	AttackAction | BurnCardAction | UseCardAction | ReRollAction | SkipRoundAction | SwitchAction | ConcedeAction
}

// AttackAction 玩家的攻击操作信息
type AttackAction struct {
	Sender uint                      `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
	Target uint                      `json:"target" yaml:"target" xml:"target"` // Target 被攻击的玩家
	Skill  uint                      `json:"skill" yaml:"skill" xml:"skill"`    // Skill 执行攻击的技能
	Paid   map[enum.ElementType]uint `json:"paid" yaml:"paid" xml:"paid"`       // Paid 支付发起攻击的费用
}

// BurnCardAction 玩家的元素转换操作信息
type BurnCardAction struct {
	Sender uint             `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
	Card   uint             `json:"card" yaml:"card" xml:"card"`       // Card 被操作的卡牌
	Paid   enum.ElementType `json:"paid" yaml:"paid" xml:"paid"`       // Paid 支付元素转换的费用
}

// ReRollAction 玩家的重掷元素骰子操作信息
type ReRollAction struct {
	Sender  uint                      `json:"sender" yaml:"sender" xml:"sender"`    // Sender 发起操作的玩家
	Dropped map[enum.ElementType]uint `json:"dropped" yaml:"dropped" xml:"dropped"` // Dropped 被舍弃的元素骰子
}

// UseCardAction 玩家的使用卡牌操作信息
type UseCardAction struct {
	Sender uint                      `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
	Card   uint                      `json:"card" yaml:"card" xml:"card"`       // Card 玩家打出的卡牌
	Paid   map[enum.ElementType]uint `json:"paid" yaml:"paid" xml:"paid"`       // Paid 玩家打出卡牌支付的费用
}

// SkipRoundAction 玩家的跳过回合操作信息
type SkipRoundAction struct {
	Sender uint `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
}

// SwitchAction 玩家的切换前台角色操作信息
type SwitchAction struct {
	Sender uint                      `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
	Target uint                      `json:"target" yaml:"target" xml:"target"` // Target 切换到的目标角色
	Paid   map[enum.ElementType]uint `json:"paid" yaml:"paid" xml:"paid"`       // Paid 玩家切换角色支付的费用
}

// ConcedeAction 玩家的弃权操作信息
type ConcedeAction struct {
	Sender uint `json:"sender" yaml:"sender" xml:"sender"` // Sender 发起操作的玩家
}

// ActionMessage 玩家的操作信息
type ActionMessage struct {
	Type   enum.ActionType `json:"type" yaml:"type" xml:"type"`       // Type 玩家的操作类型
	Sender uint            `json:"sender" yaml:"sender" xml:"sender"` // Sender 发送操作的玩家
	Args   interface{}     `json:"args" yaml:"args" xml:"args"`       // Args 玩家的操作信息
}

func exchangeMessage[target ActionMessageInterface](message ActionMessage, conditions ...enum.ActionType) (success bool, result target) {
	typeCheck := false
	for _, condition := range conditions {
		if message.Type == condition {
			typeCheck = true
			break
		}
	}

	if !typeCheck {
		return false, result
	} else {
		result, success = message.Args.(target)
		return success, result
	}
}

// ToAttackMessage 将ActionMessage转换为AttackAction
func (a ActionMessage) ToAttackMessage() (success bool, message AttackAction) {
	return exchangeMessage[AttackAction](a, enum.ActionNormalAttack, enum.ActionElementalSkill, enum.ActionElementalBurst)
}

// ToBurnCardMessage 将ActionMessage转换为BurnCardAction
func (a ActionMessage) ToBurnCardMessage() (success bool, message BurnCardAction) {
	return exchangeMessage[BurnCardAction](a, enum.ActionBurnCard)
}

// ToUesCardMessage 将ActionMessage转换为UseCardAction
func (a ActionMessage) ToUesCardMessage() (success bool, message UseCardAction) {
	return exchangeMessage[UseCardAction](a, enum.ActionUseCard)
}

// ToReRollMessage 将ActionMessage转换为ReRollAction
func (a ActionMessage) ToReRollMessage() (success bool, message ReRollAction) {
	return exchangeMessage[ReRollAction](a, enum.ActionReRoll)
}

// ToSwitchMessage 将ActionMessage转换为SwitchAction
func (a ActionMessage) ToSwitchMessage() (success bool, message SwitchAction) {
	return exchangeMessage[SwitchAction](a, enum.ActionSwitch)
}

// ToConcedeMessage 将ActionMessage转换为ConcedeAction
func (a ActionMessage) ToConcedeMessage() (success bool, message ConcedeAction) {
	return exchangeMessage[ConcedeAction](a, enum.ActionConcede)
}

// ToSkipRoundMessage 将ActionMessage转换为SkipRoundAction
func (a ActionMessage) ToSkipRoundMessage() (success bool, message SkipRoundAction) {
	return exchangeMessage[SkipRoundAction](a, enum.ActionSkipRound)
}
