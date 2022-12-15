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

type EventReactionTypeCalculatorInterface interface {
	Type() definition.RuleType
	Calculate(elementNew definition.Element, elementAttached []definition.Element) (reaction definition.Reaction, elementSurplus []definition.Element)
	ContainsRelativeReaction(elementNew definition.Element, elementAttached []definition.Element, relativeElement definition.Element) bool
}

type EventReactionDamageCalculatorInterface interface {
	Type() definition.RuleType
	Calculate(target *Player, ctx *AttackDamageContext)
}

type EventReactionEffectHandlerInterface interface {
	Type() definition.RuleType
	Handler(target *Player, ctx *DefenceDamageContext)
}

type EventRollStageHandlerInterface interface {
	Type() definition.RuleType
	Roll(setCaps uint) (set definition.ElementSet)
	ReRoll(originSet definition.ElementSet, dropSet definition.ElementSet) (result definition.ElementSet)
}
