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
	Cards() []Card
	Characters() map[uint]Character
	SetUID(uint)
	SetCards([]Card)
	SetCharacters([]Character)
}

type Player interface {
	UID() uint
	Status() enum.PlayerStatus
	Operated() bool
	ReRollTimes() uint
	StaticCost() Cost
	HoldingCost() Cost
	SwitchNextCharacter()
	SwitchPrevCharacter()
	SwitchCharacter(target uint)
	ExecuteCallbackModify(ctx *context.CallbackContext)
	ExecuteElementAttachment(ctx *context.DamageContext)
	ExecuteAttack(skill, target uint, background []uint) (result *context.DamageContext)
	ExecuteDefence(ctx *context.DamageContext)
	ExecuteModify(ctx *context.ModifierContext)
	ExecuteCharge(ctx *context.ChargeContext)
	ExecuteHeal(ctx *context.HealContext)
	ExecuteElementPayment(basic, pay Cost) (success bool)
	ExecuteElementObtain(obtain Cost)
	ExecuteElementReRoll(drop Cost)
	ExecuteBurnCard(card uint, exchangeElement enum.ElementType)
	ExecuteEatFood(card, targetCharacter uint)
	ExecuteDirectAttackModifiers(ctx *context.DamageContext)
	ExecuteFinalAttackModifiers(ctx *context.DamageContext)
	ExecuteAfterAttackCallback()
	ExecuteAfterDefenceCallback()
	ExecuteResetCallback()
	ExecuteRoundEndCallback()
	ExecuteRoundStartCallback()
	ExecuteSummonSkills()
	ExecuteAddSummonRounds(summon uint, rounds uint)
	ExecuteRemoveSummon(summon uint)
	ExecuteRemoveAllSummons()
	ExecuteSkipRound()
	ExecuteConcede()

	PreviewElementCost(basic Cost) (result Cost)

	ResetOperated()
	SetHoldingCost(cost Cost)

	GetActiveCharacter() (has bool, character Character)
	GetCharacter(id uint) (has bool, character Character)
	GetBackgroundCharacters() (characters []Character)
	HeldCard(card uint) (held bool)
	Defeated() bool
}

type player struct {
	uid    uint              // uid 玩家的UID，由其他模块托管
	status enum.PlayerStatus // status 玩家的状态

	operated    bool // operated 本回合玩家是否操作过
	reRollTimes uint // reRollTimes 重新投掷的次数
	staticCost  Cost // staticCost 每回合投掷阶段固定产出的骰子

	holdingCost     Cost                           // holdingCost 玩家持有的骰子
	cardDeck        CardDeck                       // cardDeck 玩家的牌堆
	holdingCards    kv.Map[uint, Card]             // holdingCards 玩家持有的卡牌
	activeCharacter uint                           // activeCharacter 玩家当前的前台角色
	characters      kv.OrderedMap[uint, Character] // characters 玩家出战的角色
	summons         kv.OrderedMap[uint, Summon]    // summons 玩家在场的召唤物
	supports        kv.OrderedMap[uint, Support]   // supports 玩家在场的支援

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

func (p *player) executeCallbackEvent(trigger enum.TriggerType) {
	ctx := context.NewCallbackContext()
	p.callbackEvents.Call(trigger, ctx)
	p.ExecuteCallbackModify(ctx)
}

func (p player) UID() uint {
	return p.uid
}

func (p player) Status() enum.PlayerStatus {
	return p.status
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

func (p *player) GetCharacter(id uint) (has bool, character Character) {
	return p.characters.Exists(id), p.characters.Get(id)
}

func (p *player) GetActiveCharacter() (has bool, character Character) {
	if character = p.characters.Get(p.activeCharacter); character.Status() != enum.CharacterStatusDefeated {
		return true, character
	} else {
		return false, nil
	}
}

func (p *player) GetBackgroundCharacters() (characters []Character) {
	characters = []Character{}
	p.characters.Range(func(k uint, v Character) bool {
		if k != p.activeCharacter && v.Status() != enum.CharacterStatusDefeated {
			characters = append(characters, v)
		}
		return true
	})

	return characters
}

func (p player) HeldCard(card uint) (held bool) {
	return p.holdingCards.Exists(card)
}

func (p *player) Defeated() bool {
	tag := true
	p.characters.Range(func(k uint, v Character) bool {
		if v.Status() != enum.CharacterStatusDefeated {
			tag = false
			return false
		}
		return true
	})
	return tag
}

func (p *player) ExecuteCallbackModify(ctx *context.CallbackContext) {
	// 执行ElementChangeResult
	if ok, changeElementResult := ctx.ChangeElementsResult(); ok {
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
	}

	// 执行ChargeChangeResult
	if ok, changeChargeResult := ctx.ChangeChargeResult(); ok {
		p.ExecuteCharge(changeChargeResult)
	}

	// 执行ModifiersChangeResult
	if ok, changeModifiersResult := ctx.ChangeModifiersResult(); ok {
		p.ExecuteModify(changeModifiersResult)
	}

	// 执行OperatedResult
	if switched, result := ctx.ChangeOperatedResult(); switched {
		p.operated = result
	}

	// 执行ChangeCharacter
	if switched, result := ctx.SwitchCharacterResult(); switched {
		p.SwitchCharacter(result)
	}

	// 执行GetCard
	if ok, result := ctx.GetCardsResult(); ok && result > 0 {
		for i := uint(0); i < result; i++ {
			if card, success := p.cardDeck.GetOne(); success {
				p.holdingCards.Set(card.ID(), card)
			}
		}
	}

	// 执行FindCard
	if find, target := ctx.GetFindCardResult(); find {
		if card, success := p.cardDeck.FindOne(target); success {
			p.holdingCards.Set(card.ID(), card)
		}
	}

	// 执行ElementAttachment
	if ok, attachment := ctx.AttachElementResult(); ok {
		for target, element := range attachment {
			if p.characters.Exists(target) {
				p.characters.Get(target).ExecuteElementAttachment(element)
			}
		}
	}
}

func (p *player) ExecuteElementAttachment(ctx *context.DamageContext) {
	for target, damage := range ctx.Damage() {
		p.characters.Get(target).ExecuteElementAttachment(damage.ElementType())
	}
}

func (p *player) SwitchNextCharacter() {
	// notice: 如果只有一人，此处不进行切换
	index := p.characters.GetIndex(p.activeCharacter)
	for i := index + 1; i < p.characters.Length(); i++ {
		if character := p.characters.Get(p.characters.GetKey(i)); character.Status() != enum.CharacterStatusDefeated {
			p.SwitchCharacter(character.ID())
			return
		}
	}
	for i := uint(0); i < index; i++ {
		if character := p.characters.Get(p.characters.GetKey(i)); character.Status() != enum.CharacterStatusDefeated {
			p.SwitchCharacter(character.ID())
			return
		}
	}
}

func (p *player) SwitchPrevCharacter() {
	// notice: 如果只有一人，此处不进行切换
	index := p.characters.GetIndex(p.activeCharacter)
	for i := int(index - 1); i >= 0; i-- {
		if character := p.characters.Get(p.characters.GetKey(uint(i))); character.Status() != enum.CharacterStatusDefeated {
			p.SwitchCharacter(character.ID())
			return
		}
	}
	for i := p.characters.Length(); i > index; i-- {
		if character := p.characters.Get(p.characters.GetKey(i)); character.Status() != enum.CharacterStatusDefeated {
			p.SwitchCharacter(character.ID())
			return
		}
	}
}

func (p *player) SwitchCharacter(target uint) {
	if p.characters.Exists(target) && p.activeCharacter != target {
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

func (p *player) ExecuteDirectAttackModifiers(ctx *context.DamageContext) {
	p.globalDirectAttackModifiers.Execute(ctx)
	p.characters.Get(p.activeCharacter).ExecuteDirectAttackModifiers(ctx)
}

func (p *player) ExecuteFinalAttackModifiers(ctx *context.DamageContext) {
	p.globalFinalAttackModifiers.Execute(ctx)
	p.characters.Get(p.activeCharacter).ExecuteFinalAttackModifiers(ctx)
}

func (p *player) ExecuteDefence(ctx *context.DamageContext) {
	p.globalDefenceModifiers.Execute(ctx)
	for target := range ctx.Damage() {
		p.characters.Get(target).ExecuteDefence(ctx)
	}
}

func (p *player) ExecuteAfterAttackCallback() {
	p.executeCallbackEvent(enum.AfterAttack)
}

func (p *player) ExecuteAfterDefenceCallback() {
	p.executeCallbackEvent(enum.AfterDefence)
}

func (p *player) ExecuteResetCallback() {
	p.status = enum.PlayerStatusWaiting
	p.executeCallbackEvent(enum.AfterReset)
}

func (p *player) ExecuteRoundEndCallback() {
	p.status = enum.PlayerStatusWaiting
	p.executeCallbackEvent(enum.AfterRoundEnd)
}

func (p *player) ExecuteRoundStartCallback() {
	p.status = enum.PlayerStatusReady
	p.executeCallbackEvent(enum.AfterRoundStart)
}

func (p *player) ExecuteSkipRound() {
	p.status = enum.PlayerStatusWaiting
	p.operated = true
}

func (p *player) ExecuteConcede() {
	p.status = enum.PlayerStatusDefeated
}

func (p *player) ResetOperated() {
	p.status = enum.PlayerStatusActing
	p.operated = false
}

func (p *player) ExecuteSummonSkills() {}

func (p *player) ExecuteAddSummonRounds(summon uint, rounds uint) {}

func (p *player) ExecuteRemoveSummon(summon uint) {}

func (p *player) ExecuteRemoveAllSummons() {}

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
	// 玩家持有卡牌、被转换元素不是通用元素且被转换元素数量大于0时，才可以转换
	if p.holdingCards.Exists(card) && exchangeElement != enum.ElementCurrency && p.holdingCost.costs[exchangeElement] != 0 {
		p.holdingCost.sub(exchangeElement, 1)
		p.holdingCards.Remove(card)
		p.holdingCost.add(p.characters.Get(p.activeCharacter).Vision(), 1)
		p.executeCallbackEvent(enum.AfterBurnCard)
	}
}

func (p *player) ExecuteEatFood(card, targetCharacter uint) {
	if p.holdingCards.Exists(card) && p.holdingCards.Get(card).Type() == enum.CardFood && p.characters.Exists(targetCharacter) {
		if food, ok := p.holdingCards.Get(card).(FoodCard); ok {
			ctx := context.NewModifierContext()
			food.ExecuteModify(ctx)
			p.characters.Get(targetCharacter).ExecuteEatFood(ctx)
			p.executeCallbackEvent(enum.AfterEatFood)
		}
	}
}

func (p *player) SetHoldingCost(cost Cost) {
	p.holdingCost = cost
}

func NewPlayer(info PlayerInfo) Player {
	player := &player{
		uid:                         info.UID(),
		status:                      enum.PlayerStatusViewing,
		operated:                    false,
		reRollTimes:                 1,
		staticCost:                  *NewCost(),
		holdingCost:                 *NewCost(),
		cardDeck:                    *NewCardDeck(info.Cards()),
		holdingCards:                kv.NewSimpleMap[Card](),
		activeCharacter:             0,
		characters:                  kv.NewOrderedMap[uint, Character](),
		summons:                     kv.NewOrderedMap[uint, Summon](),
		supports:                    kv.NewOrderedMap[uint, Support](),
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
