package enum

// PlayerStatus 玩家的状态
type PlayerStatus byte

const (
	PlayerStatusInitialized PlayerStatus = iota // PlayerStatusInitialized 玩家初始化完成
	PlayerStatusReady                           // PlayerStatusReady 玩家在回合开始阶段准备
	PlayerStatusWaiting                         // PlayerStatusWaiting 玩家在等待其他玩家操作
	PlayerStatusActing                          // PlayerStatusActing 玩家正在执行操作
	PlayerStatusDefeated                        // PlayerStatusDefeated 玩家被击败
	PlayerStatusViewing                         // PlayerStatusViewing 玩家正在观战
)
