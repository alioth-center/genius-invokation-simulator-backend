package entity

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

// filter 玩家行动过滤器
type filter = map[enum.ActionType]bool

var (
	cachedFilter = map[enum.PlayerStatus]map[enum.ActionType]bool{
		enum.PlayerStatusInitialized: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         true,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         false,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        true,
		},
		enum.PlayerStatusReady: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         false,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         true,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        true,
		},
		enum.PlayerStatusWaiting: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         false,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         false,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        true,
		},
		enum.PlayerStatusActing: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   true,
			enum.ActionElementalSkill: true,
			enum.ActionElementalBurst: true,
			enum.ActionSwitch:         true,
			enum.ActionBurnCard:       true,
			enum.ActionUseCard:        true,
			enum.ActionReRoll:         true,
			enum.ActionSkipRound:      true,
			enum.ActionConcede:        true,
		},
		enum.PlayerStatusDefeated: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         false,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         false,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        false,
		},
		enum.PlayerStatusViewing: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         false,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         false,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        false,
		},
		enum.PlayerStatusCharacterDefeated: {
			enum.ActionNone:           false,
			enum.ActionNormalAttack:   false,
			enum.ActionElementalSkill: false,
			enum.ActionElementalBurst: false,
			enum.ActionSwitch:         true,
			enum.ActionBurnCard:       false,
			enum.ActionUseCard:        false,
			enum.ActionReRoll:         false,
			enum.ActionSkipRound:      false,
			enum.ActionConcede:        true,
		},
	}
)

// playerChain 玩家的行动顺序表
type playerChain struct {
	canOperated map[uint]bool
	queue       []uint
	offset      int
}

// next 寻找下一个可以执行操作的玩家
func (pc *playerChain) next() (has bool, player uint) {
	for i := pc.offset; i < len(pc.queue); i++ {
		if pc.canOperated[pc.queue[i]] {
			pc.offset = i + 1
			return true, pc.queue[i]
		}
	}
	for i := 0; i < pc.offset; i++ {
		if pc.canOperated[pc.queue[i]] {
			pc.offset = i + 1
			return true, pc.queue[i]
		}
	}

	return false, 0
}

// complete 将player设置为不可执行操作
func (pc *playerChain) complete(player uint) {
	if _, exist := pc.canOperated[player]; exist {
		pc.canOperated[player] = false
	}
}

// empty 将队列清空，为复用准备
func (pc *playerChain) empty() {
	pc.canOperated = map[uint]bool{}
	pc.queue = []uint{}
	pc.offset = 0
}

// add 向队列中加入一个玩家，并将其可执行状态设置为true
func (pc *playerChain) add(player uint) {
	if _, exist := pc.canOperated[player]; !exist {
		pc.queue = append(pc.queue, player)
		pc.canOperated[player] = true
	}
}

// allActive 将队列中所有可执行状态为true的玩家导出
func (pc *playerChain) allActive() []uint {
	result := make([]uint, 0)
	for _, id := range pc.queue {
		if pc.canOperated[id] {
			result = append(result, id)
		}
	}

	return result
}

func newPlayerChain() *playerChain {
	return &playerChain{
		canOperated: map[uint]bool{},
		queue:       []uint{},
		offset:      0,
	}
}

type SyncDefeatedCharacterMessage struct {
	DefeatedPlayerUID uint
}

type SyncSwitchedCharacterMessage struct {
	SwitchedPlayerUID uint
}

type Core struct {
	players      map[uint]*player
	filters      map[uint]filter
	entities     map[uint]uint
	actingPlayer uint
	roundCount   uint
	ruleSet      RuleSet
	activeChain  *playerChain
	nextChain    *playerChain
	defeatedChan chan SyncDefeatedCharacterMessage // defeatedChan 有玩家的前台角色被击败了，需要切换角色时，会往此管道写消息
	switchedChan chan SyncSwitchedCharacterMessage // switchedChan 玩家的前台角色被击败后，切换被击败角色完成时，会往此管道写信息
	operatedChan chan struct{}                     // operatedChan 玩家结束操作时，会往此管道些信息
	errChan      chan error
}

// updatePlayerStatusAndCoreFilter 更新玩家状态与玩家可操作列表，没有做校验，请在调用前校验player不为nil且在filter中有记录
func (c *Core) updatePlayerStatusAndCoreFilter(player *player, status enum.PlayerStatus) {
	player.status = status
	c.filters[player.uid] = cachedFilter[status]
}

// filterUpdater 负责更新玩家可执行操作的协程，响应operatedChan和defeatedChan
func (c *Core) filterUpdater() {
	for existNextPlayer, actingPlayer := c.activeChain.next(); existNextPlayer; existNextPlayer, actingPlayer = c.activeChain.next() {
		if player, existPlayer := c.players[actingPlayer]; existPlayer {
			// 更新下一个玩家的状态为执行中
			c.updatePlayerStatusAndCoreFilter(player, enum.PlayerStatusActing)
			c.actingPlayer = actingPlayer

		block:
			// 阻塞更新协程并等待同步
			for {
				select {
				case defeated := <-c.defeatedChan:
					// 让当前执行的玩家进行等待
					waitingPlayer, _ := c.players[c.actingPlayer]
					c.updatePlayerStatusAndCoreFilter(waitingPlayer, enum.PlayerStatusWaiting)

					// 让被击败的玩家切换被击败的角色
					defeatedPlayer, _ := c.players[defeated.DefeatedPlayerUID]
					c.updatePlayerStatusAndCoreFilter(defeatedPlayer, enum.PlayerStatusCharacterDefeated)
				case switched := <-c.switchedChan:
					// 让切换完毕的玩家进入等待
					switchedPlayer, _ := c.players[switched.SwitchedPlayerUID]
					c.updatePlayerStatusAndCoreFilter(switchedPlayer, enum.PlayerStatusWaiting)

					// 让当前行动的玩家继续行动
					continuePlayer, _ := c.players[c.actingPlayer]
					c.updatePlayerStatusAndCoreFilter(continuePlayer, enum.PlayerStatusActing)
				case <-c.operatedChan:
					// 让当前行动的玩家进入等待
					completedPlayer, _ := c.players[c.actingPlayer]
					c.updatePlayerStatusAndCoreFilter(completedPlayer, enum.PlayerStatusWaiting)
					// 当前玩家操作完毕，退出阻塞
					break block
				}
			}
		} else {
			// 下一个玩家没有被框架托管，理论上不可能，致命错误
			c.errChan <- fmt.Errorf("error occurred while handling player %v, does not exist", actingPlayer)
			// todo: 关闭Core服务
		}
	}
}
