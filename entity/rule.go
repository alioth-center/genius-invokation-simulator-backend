package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type ReactionCalculator interface {
	// ReactionCalculate 计算当前角色身上附着的元素之间能否发生反应，返回反应类型和剩余元素
	ReactionCalculate([]enum.ElementType) (reaction enum.Reaction, elementRemains []enum.ElementType)

	// DamageCalculate 根据反应类型计算对应的伤害修正
	DamageCalculate(reaction enum.Reaction, targetCharacter uint, ctx *context.DamageContext)

	// EffectCalculate 根据反应类型计算对应的反应效果
	EffectCalculate(reaction enum.Reaction) (ctx *context.CallbackContext)

	// Attach 尝试让新元素附着在现有元素集合内，此时不触发元素反应，返回尝试附着后的元素集合
	Attach(originalElements []enum.ElementType, newElement enum.ElementType) (resultElements []enum.ElementType)

	// Relative 判断某种反应是否是某元素的相关反应
	Relative(reaction enum.Reaction, relativeElement enum.ElementType) bool
}

var (
	nullReactionCalculator ReactionCalculator = nil
)

type RuleSet interface {
	ImplementationCheck() bool
	ReactionCalculator() ReactionCalculator
}

type ruleSet struct {
	reactionCalculator ReactionCalculator
}

func (r ruleSet) ImplementationCheck() bool {
	return r.reactionCalculator != nullReactionCalculator
}

func (r ruleSet) ReactionCalculator() ReactionCalculator {
	return r.reactionCalculator
}

func NewEmptyRuleSet() RuleSet {
	return ruleSet{
		reactionCalculator: nil,
	}
}

func NewRuleSet(
	elementalReactionCalculator ReactionCalculator,
) RuleSet {
	return ruleSet{
		reactionCalculator: elementalReactionCalculator,
	}
}
