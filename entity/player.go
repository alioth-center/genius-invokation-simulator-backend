package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier/definition"
)

type PlayerInfo struct {
	UID        uint64
	Cards      []model.Card
	Characters []*character
}

type player struct {
	uid    uint64            // uid 玩家的UID，由其他模块托管
	status enum.PlayerStatus // status 玩家的状态

	operated    bool        // operated 本回合玩家是否操作过
	reRollTimes uint        // reRollTimes 重新投掷的次数
	staticCost  *model.Cost // staticCost 每回合投掷阶段固定产出的骰子

	holdingCost     *model.Cost           // holdingCost 玩家持有的骰子
	cardDeck        *CardDeck             // cardDeck 玩家的牌堆
	holdingCards    map[uint64]model.Card // holdingCards 玩家持有的卡牌
	activeCharacter uint64                // activeCharacter 玩家当前的前台角色

	characters    map[uint64]*character // characters 玩家出战的角色
	characterList []uint64              // characterList 玩家的角色列表
	summons       map[uint64]Summon     // summons 玩家在场的召唤物
	summonList    []uint64              // summonList 玩家的召唤物列表
	supports      map[uint64]Support    // supports 玩家在场的支援
	supportList   []uint64              // supportList 玩家的支援物列表

	globalDirectAttackModifiers definition.AttackModifiers  // globalDirectAttackModifiers 全局直接攻击修正
	globalFinalAttackModifiers  definition.AttackModifiers  // globalFinalAttackModifiers 全局最终攻击修正
	globalDefenceModifiers      definition.DefenceModifiers // globalDefenceModifiers 全局防御修正
	globalHealModifiers         definition.HealModifiers    // globalHealModifiers 全局治疗修正
	globalChargeModifiers       definition.ChargeModifiers  // globalChargeModifiers 全局充能修正
	globalCostModifiers         definition.CostModifiers    // globalCostModifiers 全局费用修正

	cooperativeAttacks map[enum.TriggerType]model.CooperativeSkill // cooperativeAttacks 协同攻击技能
	callbackEvents     *Map                                        // callbackEvents 回调事件集合
}

func (p player) GetUID() (uid uint64) {
	return p.uid
}

func (p player) GetCost() (cost map[enum.ElementType]uint) {
	return p.holdingCost.Costs()
}

func (p player) GetCards() (cards []uint64) {
	cards = []uint64{}
	for i := range p.holdingCards {
		cards = append(cards, i)
	}

	return cards
}

func (p player) GetSummons() (summons []uint64) {
	return p.summonList
}

func (p player) GetSupports() (supports []uint64) {
	return p.summonList
}

func (p player) CardDeckRemain() (remain uint) {
	return p.cardDeck.remain
}

func (p player) GetActiveCharacter() (character uint64) {
	return p.activeCharacter
}

func (p player) GetBackgroundCharacters() (characters []uint64) {
	characters = []uint64{}
	for _, character := range p.characters {
		if character.status != enum.CharacterStatusDefeated && character.id != p.activeCharacter {
			characters = append(characters, character.id)
		}
	}

	return characters
}

func (p player) GetCharacter(character uint64) (has bool, entity model.Character) {
	characterEntity, exist := p.characters[character]
	return exist, characterEntity
}

func (p player) GetStatus() (status enum.PlayerStatus) {
	return p.status
}

func (p player) GetGlobalModifiers(modifierType enum.ModifierType) (modifiers []uint64) {
	switch modifierType {
	case enum.ModifierTypeNone:
		return []uint64{}
	case enum.ModifierTypeAttack:
		modifiers = p.globalDirectAttackModifiers.Expose()
		modifiers = append(modifiers, p.globalFinalAttackModifiers.Expose()...)
		return modifiers
	case enum.ModifierTypeCharacter:
		return []uint64{}
	case enum.ModifierTypeCharge:
		return p.globalChargeModifiers.Expose()
	case enum.ModifierTypeCost:
		return p.globalCostModifiers.Expose()
	case enum.ModifierTypeDefence:
		return p.globalDefenceModifiers.Expose()
	case enum.ModifierTypeHeal:
		return p.globalHealModifiers.Expose()
	default:
		return []uint64{}
	}
}

func (p player) GetCooperativeSkills(trigger enum.TriggerType) (skills []uint64) {
	return []uint64{}
}

func (p player) GetEvents(trigger enum.TriggerType) (events []uint64) {
	return p.callbackEvents.Expose(trigger)
}
