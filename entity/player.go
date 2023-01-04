package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/event"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type PlayerInfo interface {
	UID() uint
	Name() string
	Cards()
	Characters() map[uint]Character
}

type Player interface {
	UID() uint
	Name() string
	Operated() bool
	ReRollTimes() uint
	StaticCost() Cost
	HoldingCost() Cost
	SwitchCharacter(target uint)
	ExecuteAttack(skill, target uint, background []uint) (result *context.DamageContext)
	ExecuteModify(ctx *context.ModifierContext)
	ExecuteCharge(ctx *context.ChargeContext)
	ExecuteHeal(ctx *context.HealContext)
	ExecuteElementPayment(basic, pay Cost) (success bool)
	ExecuteElementObtain(obtain Cost)
	ExecuteElementReRoll(drop Cost)
	ExecuteBurnCard(card uint, exchangeElement enum.ElementType)
}

type player struct {
	uid  uint   // uid 玩家的UID，由其他模块托管
	name string // name 玩家的名称

	operated    bool // operated 本回合玩家是否操作过
	reRollTimes uint // reRollTimes 重新投掷的次数
	staticCost  Cost // staticCost 每回合投掷阶段固定产出的骰子

	holdingCost     Cost                    // holdingCost 玩家持有的骰子
	holdingCards    kv.Map[uint, uint]      // holdingCards 玩家持有的卡牌
	activeCharacter uint                    // activeCharacter 玩家当前的前台角色
	characters      kv.Map[uint, Character] // characters 玩家出战的角色
	summons         kv.Map[uint, Summon]    // summons 玩家在场的召唤物
	supports        kv.Map[uint, Support]   // supports 玩家在场的支援

	globalDirectAttackModifiers AttackModifiers  // globalDirectAttackModifiers 全局直接攻击修正
	globalFinalAttackModifiers  AttackModifiers  // globalFinalAttackModifiers 全局最终攻击修正
	globalDefenceModifiers      DefenceModifiers // globalDefenceModifiers 全局防御修正
	globalHealModifiers         HealModifiers    // globalHealModifiers 全局治疗修正
	globalChargeModifiers       ChargeModifiers  // globalChargeModifiers 全局充能修正
	globalCostModifiers         CostModifiers    // globalCostModifiers 全局费用修正

	cooperativeAttacks []CooperativeSkill // cooperativeAttacks 协同攻击技能
	callbackEvents     event.Map          // callbackEvents 回调事件集合
}

func (p *player) executeCharacterModify(ctx *context.ModifierContext) {
	p.characters.Range(func(id uint, character Character) bool {
		character.ExecuteModify(ctx)
		return true
	})
}

func (p *player) executeCallbackModify(ctx *context.CallbackContext) {
	// 执行ElementChangeResult
	changeElementResult := ctx.ChangeElementsResult()
	addElement, removeElement := map[enum.ElementType]uint{}, map[enum.ElementType]uint{}
	for element, amount := range changeElementResult.Cost() {
		if amount > 0 {
			addElement[element] += uint(amount)
		} else {
			removeElement[element] += uint(-amount)
		}
	}
	p.holdingCost.Pay(*NewCostFromMap(removeElement))
	p.holdingCost.Add(*NewCostFromMap(addElement))

	// 执行ChargeChangeResult
	changeChargeResult := ctx.ChangeChargeResult()
	p.ExecuteCharge(&changeChargeResult)

	// 执行ModifiersChangeResult
	changeModifiersResult := ctx.ChangeModifiersResult()
	p.ExecuteModify(&changeModifiersResult)

	// 执行OperatedResult
	if switched, result := ctx.ChangeOperatedResult(); switched {
		p.operated = result
	}

	// 执行ChangeCharacter
	if switched, result := ctx.SwitchCharacterResult(); switched {
		p.SwitchCharacter(result)
	}

	// 执行GetCard
	if ctx.GetCardsResult() > 0 {
		// todo: implement me
		panic("implement card-deck")
	}

	// 执行ElementAttachment
	attachment := ctx.AttachElementResult()
	for target := range attachment {
		if p.characters.Exists(target) {
			// todo: implement me
			panic("implement rule.reaction")
		}
	}

}

func (p *player) executeCallbackEvent(trigger enum.TriggerType) {
	ctx := context.NewCallbackContext()
	p.callbackEvents.Call(trigger, ctx)
	p.executeCallbackModify(ctx)
}

func (p player) UID() uint {
	return p.uid
}

func (p player) Name() string {
	return p.name
}

func (p player) Operated() bool {
	return p.operated
}

func (p player) ReRollTimes() uint {
	return p.reRollTimes
}

func (p player) StaticCost() Cost {
	return p.staticCost
}

func (p player) HoldingCost() Cost {
	return p.holdingCost
}

func (p *player) SwitchCharacter(target uint) {
	if p.characters.Exists(target) {
		p.characters.Get(p.activeCharacter).SwitchDown()
		p.activeCharacter = target
		p.characters.Get(target).SwitchUp()
		p.executeCallbackEvent(enum.AfterSwitch)
	}
}

func (p *player) ExecuteModify(ctx *context.ModifierContext) {
	if ctx.AddGlobalChargeModifiers() != nil {
		for _, m := range ctx.AddGlobalChargeModifiers() {
			p.globalChargeModifiers.Append(m)
		}
	}

	if ctx.AddGlobalCostModifiers() != nil {
		for _, m := range ctx.AddGlobalCostModifiers() {
			p.globalCostModifiers.Append(m)
		}
	}

	if ctx.AddGlobalDefenceModifiers() != nil {
		for _, m := range ctx.AddGlobalDefenceModifiers() {
			p.globalDefenceModifiers.Append(m)
		}
	}

	if ctx.AddGlobalDirectAttackModifiers() != nil {
		for _, m := range ctx.AddGlobalDirectAttackModifiers() {
			p.globalDirectAttackModifiers.Append(m)
		}
	}

	if ctx.AddGlobalFinalAttackModifiers() != nil {
		for _, m := range ctx.AddGlobalFinalAttackModifiers() {
			p.globalFinalAttackModifiers.Append(m)
		}
	}

	if ctx.AddGlobalHealModifiers() != nil {
		for _, m := range ctx.AddGlobalHealModifiers() {
			p.globalHealModifiers.Append(m)
		}
	}

	if ctx.RemoveGlobalChargeModifiers() != nil {
		for _, m := range ctx.RemoveGlobalChargeModifiers() {
			p.globalChargeModifiers.Remove(m.ID())
		}
	}

	if ctx.RemoveGlobalCostModifiers() != nil {
		for _, m := range ctx.RemoveGlobalCostModifiers() {
			p.globalCostModifiers.Remove(m.ID())
		}
	}

	if ctx.RemoveGlobalDefenceModifiers() != nil {
		for _, m := range ctx.RemoveGlobalDefenceModifiers() {
			p.globalDefenceModifiers.Remove(m.ID())
		}
	}

	if ctx.RemoveGlobalDirectAttackModifiers() != nil {
		for _, m := range ctx.RemoveGlobalDirectAttackModifiers() {
			p.globalDirectAttackModifiers.Remove(m.ID())
		}
	}

	if ctx.RemoveGlobalFinalAttackModifiers() != nil {
		for _, m := range ctx.RemoveGlobalFinalAttackModifiers() {
			p.globalFinalAttackModifiers.Remove(m.ID())
		}
	}

	if ctx.RemoveGlobalHealModifiers() != nil {
		for _, m := range ctx.RemoveGlobalHealModifiers() {
			p.globalHealModifiers.Remove(m.ID())
		}
	}

	p.executeCharacterModify(ctx)
}

func (p *player) ExecuteCharge(ctx *context.ChargeContext) {
	p.globalChargeModifiers.Execute(ctx)
	for characterID := range ctx.Charge() {
		p.characters.Get(characterID).ExecuteCharge(ctx)
	}
	p.executeCallbackEvent(enum.AfterCharge)
}

func (p *player) ExecuteHeal(ctx *context.HealContext) {
	p.globalHealModifiers.Execute(ctx)
	for characterID := range ctx.Heal() {
		p.characters.Get(characterID).ExecuteHeal(ctx)
	}
	p.executeCallbackEvent(enum.AfterHeal)
}

func (p *player) PreviewElementCost(basic Cost) (result Cost) {
	ctx := context.NewCostContext()
	p.globalCostModifiers.Preview(ctx)
	p.characters.Get(p.activeCharacter).PreviewCostModify(ctx)
	for element, amount := range ctx.Cost() {
		if amount < 0 {
			if basic.costs[element] > uint(-amount) {
				basic.costs[element] -= uint(-amount)
			} else {
				basic.costs[element] = 0
			}
		} else {
			basic.costs[element] += uint(amount)
		}
	}

	return basic
}

func (p *player) ExecuteAttack(skill, target uint, background []uint) (result *context.DamageContext) {
	return p.characters.Get(p.activeCharacter).ExecuteAttack(skill, target, background)
}

func (p *player) ExecuteElementPayment(basic, pay Cost) (success bool) {
	if p.PreviewElementCost(basic).Equals(pay) {
		ctx := context.NewCostContext()
		p.globalCostModifiers.Execute(ctx)
		p.characters.Get(p.activeCharacter).ExecuteCostModify(ctx)
		p.holdingCost.Pay(pay)
		return true
	} else {
		return false
	}
}

func (p *player) ExecuteElementObtain(obtain Cost) {
	p.holdingCost.Add(obtain)
}

func (p *player) ExecuteElementReRoll(drop Cost) {
	p.holdingCost.Pay(drop)
	result := NewRandomCost(drop.total)
	p.holdingCost.Add(*result)
}

func (p *player) ExecuteBurnCard(card uint, exchangeElement enum.ElementType) {
	if p.holdingCards.Get(card) != 0 && p.holdingCost.costs[exchangeElement] != 0 {
		p.holdingCost.sub(exchangeElement, 1)
		p.holdingCards.Set(card, p.holdingCards.Get(card)-1)
		p.holdingCost.add(p.characters.Get(p.activeCharacter).Vision(), 1)
	}
}

func NewPlayer(info PlayerInfo) Player {
	player := &player{
		uid:                         info.UID(),
		name:                        info.Name(),
		operated:                    false,
		reRollTimes:                 1,
		staticCost:                  *NewCost(),
		holdingCost:                 *NewCost(),
		holdingCards:                kv.NewSimpleMap[uint](),
		activeCharacter:             0,
		characters:                  kv.NewSimpleMap[Character](),
		summons:                     kv.NewSimpleMap[Summon](),
		supports:                    kv.NewSimpleMap[Support](),
		globalDirectAttackModifiers: modifier.NewChain[context.DamageContext](),
		globalFinalAttackModifiers:  modifier.NewChain[context.DamageContext](),
		globalDefenceModifiers:      modifier.NewChain[context.DamageContext](),
		globalHealModifiers:         modifier.NewChain[context.HealContext](),
		globalChargeModifiers:       modifier.NewChain[context.ChargeContext](),
		globalCostModifiers:         modifier.NewChain[context.CostContext](),
		cooperativeAttacks:          nil,
		callbackEvents:              *event.NewEventMap(),
	}

	for id, character := range info.Characters() {
		if player.activeCharacter == 0 {
			player.activeCharacter = id
		}

		player.characters.Set(id, character)
	}

	return player
}
