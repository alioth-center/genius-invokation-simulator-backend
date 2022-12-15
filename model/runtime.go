/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "runtime.go" LastUpdatedAt 2022/12/12 15:16:12
 */

package model

import (
	"fmt"
)

type RuleSet struct {
	Rules map[string]RuleInterface
}

type RuntimeRuleSet struct {
	rules map[string]RuleInterface
}

type Rule struct {
	Name      string
	Implement RuleInterface
}

func initRuntimeRuleSet(rules ...Rule) *RuntimeRuleSet {
	nrs := &RuntimeRuleSet{
		rules: map[string]RuleInterface{},
	}

	for _, rule := range rules {
		nrs.rules[rule.Name] = rule.Implement
	}

	return nrs
}

var (
	ShufflePlayerChainFunction       EventShufflePlayerChainInterface       = nil // ShufflePlayerChainFunction 确定玩家初始顺序的规则实现
	ShuffleCardStackFunction         EventShuffleCardStackInterface         = nil // ShuffleCardStackFunction 将牌堆打断的规则实现
	RollStageHandlerFunction         EventRollStageHandlerInterface         = nil // RollStageHandlerFunction 投掷阶段计算骰子的规则实现
	ReactionCalculatorFunction       EventReactionTypeCalculatorInterface   = nil // ReactionCalculatorFunction 计算反应类型的规则实现
	ReactionDamageCalculatorFunction EventReactionDamageCalculatorInterface = nil // ReactionDamageCalculatorFunction 计算反应伤害的规则实现
	ReactionEffectHandlerFunction    EventReactionEffectHandlerInterface    = nil // ReactionEffectHandlerFunction 执行反应附加效果的规则实现

	RuntimeRules RuntimeRuleSet = *initRuntimeRuleSet(
		Rule{Name: "ShufflePlayerChain", Implement: ShufflePlayerChainFunction},
		Rule{Name: "ShuffleCardStack", Implement: ShuffleCardStackFunction},
		Rule{Name: "RollStageHandler", Implement: RollStageHandlerFunction},
		Rule{Name: "ReactionTypeCalculatorFunction", Implement: ReactionCalculatorFunction},
		Rule{Name: "ReactionDamageCalculatorFunction", Implement: ReactionDamageCalculatorFunction},
	)
)

func InjectRuntimeRules(rules RuleSet) {
	for name := range RuntimeRules.rules {
		if rule, ok := rules.Rules[name]; rule == nil || !ok {
			panic(fmt.Sprintf("missing necessary function: %s", name))
		}
	}

	for name := range RuntimeRules.rules {
		RuntimeRules.rules[name] = rules.Rules[name]
	}
}

func CheckImplementation() (err error) {
	for name, implementation := range RuntimeRules.rules {
		if implementation == nil {
			return fmt.Errorf("missing necessary function: %s", name)
		}
	}

	return nil
}
