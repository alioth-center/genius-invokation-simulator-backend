package definition

// Action 玩家可进行的操作类型
type Action uint

const (
	ActionNormalAttack   Action = iota // ActionNormalAttack 进行普通攻击
	ActionElementalSkill               // ActionElementalSkill 释放元素战技
	ActionElementalBurst               // ActionElementalBurst 释放元素爆发
	ActionSwitch                       // ActionSwitch 切换在场角色
	ActionBurnCard                     // ActionBurnCard 将卡牌转换为元素
	ActionUseCard                      // ActionUseCard 使用卡牌
	ActionSkipRound                    // ActionSkipRound 结束回合
	ActionConcede                      // ActionConcede 弃权让步
)
