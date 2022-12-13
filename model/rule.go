/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "rule.go" LastUpdatedAt 2022/12/12 16:37:12
 */

package model

import "github.com/sunist-c/genius-invokation-simulator-backend/definition"

// RuleInterface 规则接口
type RuleInterface interface {
	Type() definition.RuleType
}

// EventShuffleCardStackInterface 将牌堆打乱的接口
type EventShuffleCardStackInterface interface {
	Type() definition.RuleType
	Shuffle(start, end int, array []ICard)
}

// EventShufflePlayerChainInterface 将玩家行动位置打乱的接口
type EventShufflePlayerChainInterface interface {
	Type() definition.RuleType
	Shuffle(start, end int, array []*Player)
}

type EventReactionCalculatorInterface interface {
	Type() definition.RuleType
	Calculate(elementNew definition.Element, elementAttached []definition.Element) (reaction definition.Reaction, elementSurplus []definition.Element)
}

type EventReactionHandlerInterface interface {
	Type() definition.RuleType
	Handle(ctx *Context)
}

type EventRollStageHandlerInterface interface {
	Type() definition.RuleType
	Roll(setCaps uint) (set definition.ElementSet)
	ReRoll(originSet definition.ElementSet, dropSet definition.ElementSet) (result definition.ElementSet)
}

type EventNormalAttackInterface interface {
	Type() definition.RuleType
	Handle(ctx *Context, targetPlayer *Player, targetCharacter *Character, element definition.Element, skill ISkill, beforeEvents, afterEvents []IEffect)
}

type EventElementalSkillInterface interface{}

type EventElementalBurstInterface interface{}

type EventPassiveSkillInterface interface{}

type EventSwitchCharacterInterface interface{}

type EventStartGameInterface interface{}

type EventOnHitInterface interface{}

type EventOnSupportInterface interface{}

type EventOnSummonInterface interface{}

type EventOnEquipInterface interface{}

type EventOnReactionInterface interface{}

type EventOnBurnCardInterface interface{}
