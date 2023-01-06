package enum

// TriggerType 触发器类型
type TriggerType byte

const (
	AfterAttack     TriggerType = iota // AfterAttack 攻击结算后触发
	AfterBurnCard                      // AfterBurnCard 使用卡牌转换元素后触发
	AfterCharge                        // AfterCharge 充能后触发
	AfterDefence                       // AfterDefence 防御结算后触发
	AfterEatFood                       // AfterEatFood 使用食物后触发
	AfterHeal                          // AfterHeal 治疗后触发
	AfterReset                         // AfterReset 重置时触发
	AfterRoundStart                    // AfterRoundStart 回合开始后触发
	AfterRoundEnd                      // AfterRoundEnd 回合结束后触发
	AfterSwitch                        // AfterSwitch 切换角色后触发
)
