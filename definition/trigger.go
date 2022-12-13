/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "trigger.go" LastUpdatedAt 2022/12/12 10:20:12
 */

package definition

// TriggerType 事件触发器，定义的事件类型enum，触发该类型的触发器以后，将会调用触发者的Effector
type TriggerType byte

const (
	TriggerNone                 TriggerType = iota // TriggerNone 一次性触发器，执行时触发
	TriggerBeforePassiveSkill                      // TriggerBeforePassiveSkill 在被动技能触发前触发
	TriggerAfterPassiveSkill                       // TriggerAfterPassiveSkill 在被动技能触发后触发
	TriggerBeforeNormalAttack                      // TriggerBeforeNormalAttack 在普通攻击前触发
	TriggerAfterNormalAttack                       // TriggerAfterNormalAttack 在普通攻击后触发
	TriggerBeforeElementalSkill                    // TriggerBeforeElementalSkill 在元素战技前触发
	TriggerAfterElementalSkill                     // TriggerAfterElementalSkill 在元素战技后触发
	TriggerBeforeElementalBurst                    // TriggerBeforeElementalBurst 在元素爆发前触发
	TriggerAfterElementBurst                       // TriggerAfterElementBurst 在元素爆发后触发
	TriggerBeforeHit                               // TriggerBeforeHit 在受到攻击前触发
	TriggerAfterHit                                // TriggerAfterHit 在受到攻击后触发
	TriggerBeforeSwitch                            // TriggerBeforeSwitch 在切换角色前触发
	TriggerAfterSwitch                             // TriggerAfterSwitch 在切换角色后触发
	TriggerBeforeSupport                           // TriggerBeforeSupport 在执行支援前触发
	TriggerAfterSupport                            // TriggerAfterSupport 在执行支援后触发
	TriggerBeforeSummon                            // TriggerBeforeSummon 在召唤前触发
	TriggerAfterSummon                             // TriggerAfterSummon 在召唤后触发
	TriggerBeforeEquip                             // TriggerBeforeEquip 在装备前触发
	TriggerAfterEquip                              // TriggerAfterEquip 在装备后触发
	TriggerBeforeReaction                          // TriggerBeforeReaction 在元素反应前触发
	TriggerAfterReaction                           // TriggerAfterReaction 在元素反应后触发
	TriggerBeforeBurnCard                          // TriggerBeforeBurnCard 在使用卡牌转换元素前触发
	TriggerAfterBurnCard                           // TriggerAfterBurnCard 在使用卡牌转换元素后触发
	TriggerBeforeHealing                           // TriggerBeforeHealing 在角色接受治疗前触发
	TriggerAfterHealing                            // TriggerAfterHealing 在角色接受治疗后触发
	TriggerAfterGameStart                          // TriggerAfterGameStart 在对局开始后触发
	TriggerAfterRollStage                          // TriggerAfterRollStage 投掷阶段结束后触发
	TriggerAfterRoundStart                         // TriggerAfterRoundStart 回合开始后触发
	TriggerAfterRoundBattle                        // TriggerAfterRoundBattle 在回合战斗后触发
	TriggerAfterRoundSummon                        // TriggerAfterRoundSummon 在召唤物行动结束后触发
	TriggerAfterRoundSupport                       // TriggerAfterRoundSupport 在支援物行动结束后触发
	TriggerAfterRoundEnd                           // TriggerAfterRoundEnd 在回合结束后触发
)
