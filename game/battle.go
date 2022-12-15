/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "battle.go" LastUpdatedAt 2022/12/15 10:24:15
 */

package game

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"
)

type BattleFramework struct {
}

// PreviewAttack 预览本次攻击的最终伤害效果，不包含协同攻击
func (b BattleFramework) PreviewAttack(sender, target *model.Player, skill model.IAttackSkill) *model.DefenceDamageContext {
	damage := sender.ActiveCharacter.Clone().Attack(target, skill)
	model.ReactionDamageCalculatorFunction.Calculate(target, damage)
	return target.ActiveCharacter.Clone().Defense(sender, damage)
}

// ExecuteAttack 执行攻击操作
func (b BattleFramework) ExecuteAttack(sender, target *model.Player, skill model.IAttackSkill) {
	damage := sender.ActiveCharacter.Attack(target, skill)
	model.ReactionDamageCalculatorFunction.Calculate(target, damage)
	result := target.ActiveCharacter.Defense(sender, damage)
	if result.Effective() {
		for executeCharacter, executeDamage := range result.Damage() {
			if executeCharacter.Status != definition.CharacterStatusUnselectable {
				if executeCharacter.CurrentHealthPoint > executeDamage.Amount {
					executeCharacter.CurrentHealthPoint -= executeDamage.Amount
				} else {
					executeCharacter.CurrentHealthPoint = 0
					executeCharacter.Status = definition.CharacterStatusUnselectable
				}
			}
		}
	}
	model.ReactionEffectHandlerFunction.Handler(target, result)
}

// PreviewHeal 预览治疗操作的最终执行效果
func (b BattleFramework) PreviewHeal(target *model.Character, amount uint) *model.HealContext {
	return target.Clone().Heal(amount)
}

// ExecuteHeal 执行治疗操作
func (b BattleFramework) ExecuteHeal(target *model.Character, amount uint) {
	heal := target.Heal(amount)
	if target.CurrentHealthPoint+heal.Heal() < target.MaxHealthPoint {
		target.CurrentHealthPoint += heal.Heal()
	} else {
		target.CurrentHealthPoint = target.MaxHealthPoint
	}
}

// PreviewCost 预览某可消耗元素骰子的操作最终的耗费
func (b BattleFramework) PreviewCost(receiver *model.Character, cost model.IConsumable) *model.CostContext {
	costData := model.NewCostContext(cost.Cost())
	if receiver != nil {
		ctx := model.NewContext(receiver.CostModifiers, costData)
		ctx.Continue()
		costData = ctx.Data
	}
	return costData
}

// ExecuteCost 消耗元素骰子，如果操作非法，返回false
func (b BattleFramework) ExecuteCost(receiver *model.Character, player *model.Player, cost model.IConsumable, pay definition.ElementSet) bool {
	costData := b.PreviewCost(receiver, cost)
	if model.ElementSetContains(player.Elements, pay) && model.ElementSetEqual(costData.Cost(), pay) {
		player.Elements = model.SubElementSet(player.Elements, pay)
		return true
	} else {
		return false
	}
}
