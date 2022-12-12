/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "lifecycle.go" LastUpdatedAt 2022/12/12 16:58:12
 */

package definition

// RoundLifeCycle 回合内生命周期，空四位供权重使用
type RoundLifeCycle byte

const (
	RollStage      RoundLifeCycle = iota << 4 // RollStage 投掷阶段
	BeginStage                                // BeginStage 开始阶段
	BattleStage                               // BattleStage 战斗阶段
	SummonedStage                             // SummonedStage 召唤物行动阶段
	SupporterStage                            // SupporterStage 支援物行动阶段
	EndStage                                  // EndStage 结束阶段
)
