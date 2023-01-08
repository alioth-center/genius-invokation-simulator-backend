package enum

// RoundStage 回合生命周期
type RoundStage byte

const (
	RoundStageInitialized RoundStage = iota // RoundStageInitialized 初始化完成阶段，仅在游戏开始时进入
	RoundStageStart                         // RoundStageStart 回合开始阶段
	RoundStageRoll                          // RoundStageRoll 投掷阶段
	RoundStageBattle                        // RoundStageBattle 对战阶段
	RoundStageEnd                           // RoundStageEnd 回合结束阶段
)
