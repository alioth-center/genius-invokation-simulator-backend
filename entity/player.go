package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/event"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type Player struct {
	uid  uint   // uid 玩家的UID，由其他模块托管
	name string // name 玩家的名称

	operated    bool // operated 本回合玩家是否操作过
	reRollTimes uint // reRollTimes 重新投掷的次数
	staticCost  Cost // staticCost 每回合投掷阶段固定产出的骰子

	holdingCost Cost                    // holdingCost 玩家持有的骰子
	characters  kv.Map[uint, Character] // characters 玩家出站的角色
	summons     kv.Map[uint, Summon]    // summons 玩家在场的召唤物
	supports    kv.Map[uint, Support]   // supports 玩家在场的支援

	globalDirectAttackModifiers AttackModifiers  // globalDirectAttackModifiers 全局直接攻击修正
	globalFinalAttackModifiers  AttackModifiers  // globalFinalAttackModifiers 全局最终攻击修正
	globalDefenceModifiers      DefenceModifiers // globalDefenceModifiers 全局防御修正

	cooperativeAttacks []CooperativeSkill // cooperativeAttacks 协同攻击技能
	callbackEvents     event.Set          // callbackEvents 回调事件集合
}
