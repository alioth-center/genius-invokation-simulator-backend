/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "tigger.go" LastUpdatedAt 2022/12/16 10:38:16
 */

package definition

type Trigger byte

const (
	AfterAttack  Trigger = iota // AfterAttack 攻击结算后触发
	AfterDefence                // AfterDefence 防御结算后触发
	AfterEatFood                // AfterEatFood 使用食物后触发
	AfterHeal                   // AfterHeal 治疗后触发
	AfterCharge                 // AfterCharge 充能后触发
	AfterSwitch                 // AfterSwitch 切换角色后触发
)
