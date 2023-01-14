package entity

import (
	"encoding/json"
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/event"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/websocket"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/websocket/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"sync"
)

// filter 玩家行动过滤器
type filter = map[enum.ActionType]bool

var (
	nullDirectAttackModifiers kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullFinalAttackModifiers  kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullDefenceModifiers      kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullChargeModifiers       kv.Map[uint, []modifier.Modifier[context.ChargeContext]] = nil
	nullHealModifiers         kv.Map[uint, []modifier.Modifier[context.HealContext]]   = nil
	nullCostModifiers         kv.Map[uint, []modifier.Modifier[context.CostContext]]   = nil

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

type playerContext struct {
	player     *player
	connection *websocket.Connection
	filter     filter
}

type Core struct {
	errorHandler util.ErrorHandler                 // errorHandler 处理错误的日志器
	room         map[uint]*playerContext           // room 房间信息，包括玩家、合法操作
	viewers      map[uint]*websocket.Connection    // viewers 观战者的连接集合
	guests       map[*websocket.Connection]bool    // guests 匿名观战者的连接集合
	entities     map[uint]uint                     // entities 实体表
	actingPlayer uint                              // actingPlayer 当前正在操作的玩家
	roundCount   uint                              // roundCount 回合数
	roundStage   enum.RoundStage                   // roundStage 当前回合阶段
	ruleSet      model.RuleSet                     // ruleSet 当前战斗的规则
	activeChain  *playerChain                      // activeChain 当前回合的结算队列
	nextChain    *playerChain                      // nextChain 下个回合的结算队列
	defeatedChan chan SyncDefeatedCharacterMessage // defeatedChan 有玩家的前台角色被击败了，需要切换角色时，会往此管道写消息
	switchedChan chan SyncSwitchedCharacterMessage // switchedChan 玩家的前台角色被击败后，切换被击败角色完成时，会往此管道写信息
	operatedChan chan struct{}                     // operatedChan 玩家结束操作时，会往此管道写信息
	readChan     chan message.ActionMessage        // readChan 有网络信息传入时，向此管道写入信息
	exitChan     chan struct{}                     // exitChan 当结束服务时，向此管道写入信息
	updateMutex  sync.RWMutex                      // updateMutex 更新锁，用于避免并发更新玩家状态以引发并发问题
}

// updatePlayerStatusAndCoreFilter 更新玩家状态与玩家可操作列表，没有做校验，请在调用前校验player不为nil且在filter中有记录
func (c *Core) updatePlayerStatusAndCoreFilter(player *player, status enum.PlayerStatus) {
	c.updateMutex.Lock()
	player.status = status
	c.room[player.uid].filter = cachedFilter[status]
	c.updateMutex.Unlock()
}

// messageFilter 过滤不合法的操作类型，若玩家发送的操作信息合法，则legal为真
func (c *Core) messageFilter(msg message.ActionMessage) (legal bool) {
	// 如果在执行更改操作，等待更改完成
	c.updateMutex.RLock()
	defer c.updateMutex.RUnlock()
	if playerContext, exist := c.room[msg.Sender]; exist && playerContext != nil {
		if playerContext.filter[msg.Type] {
			// 操作类型合法
			return true
		} else {
			// 操作类型不合法
			return false
		}
	} else {
		// 玩家不存在
		return false
	}
}

// handleMessage 处理玩家的信息
func (c *Core) handleMessage(msg message.ActionMessage) {
	switch msg.Type {
	case enum.ActionNormalAttack:
		// 普通攻击的逻辑
		if success, attackAction := msg.ToAttackMessage(); success {
			c.executeAttack(attackAction)
		}
	case enum.ActionElementalSkill:
		// 元素战技的逻辑
		if success, attackAction := msg.ToAttackMessage(); success {
			c.executeAttack(attackAction)
		}
	case enum.ActionElementalBurst:
		// 元素爆发的逻辑
		if success, attackAction := msg.ToAttackMessage(); success {
			c.executeAttack(attackAction)
		}
	case enum.ActionSwitch:
		// 切换角色的逻辑
		if success, switchAction := msg.ToSwitchMessage(); success {
			c.executeSwitch(switchAction)
		}
	case enum.ActionBurnCard:
		// 转换元素的逻辑
		if success, burnCardAction := msg.ToBurnCardMessage(); success {
			c.executeBurnCard(burnCardAction)
		}
	case enum.ActionUseCard:
		// 使用卡牌的逻辑
		if success, useCardAction := msg.ToUesCardMessage(); success {
			c.executeUseCard(useCardAction)
		}
	case enum.ActionReRoll:
		// 重掷骰子的逻辑
		if success, switchAction := msg.ToSwitchMessage(); success {
			c.executeSwitch(switchAction)
		}
	case enum.ActionSkipRound:
		// 跳过回合的逻辑
		if success, skipAction := msg.ToSkipRoundMessage(); success {
			c.executeSkipRound(skipAction)
		}
	case enum.ActionConcede:
		// 比赛弃权的逻辑
		if success, concedeAction := msg.ToConcedeMessage(); success {
			c.executeConcede(concedeAction)
		}
	default:
		// 没有解析到messageType，理论上不可能，当作恶意网络包拦截
		c.errorHandler.Handle(fmt.Errorf("unknown action %d, sent by %v", msg.Type, msg.Sender))
	}

	// 处理完毕，需要向所有玩家同步信息
	c.sendSyncMessage()
}

// generateSyncMessage 生成某玩家收到的同步信息，playerID为0则生成匿名访客的观战信息
func (c *Core) generateSyncMessage(player uint) (syncMessage message.SyncMessage) {
	dictionary := generateDictionary(c)
	background := generateBackgroundMessage(c)
	if c.room[player] != nil {
		// 是参与对战玩家，需要开战争迷雾
		playerEntity := c.room[player].player
		playerMessage := message.PlayerMessage{
			Self:   generateSelfMessage(c, playerEntity),
			Others: generateOtherMessage(c, playerEntity),
			Append: dictionary,
		}

		return message.NewSyncMessage(player, playerMessage, background)
	} else if player != 0 {
		if c.viewers[player] != nil {
			// 不是参与对战玩家，但是认证了，开上帝视野
			var playerList []message.Self

			for _, playerContext := range c.room {
				playerList = append(playerList, generateSelfMessage(c, playerContext.player))
			}

			viewerMessage := message.ViewerMessage{
				Players: playerList,
				Append:  dictionary,
			}

			return message.NewSyncMessage(player, viewerMessage, background)
		} else {
			// 不是参与对战玩家，且没有认证，理论上不可能，返回默认的战争迷雾视角
			playerList := generateOtherMessage(c, nil)

			guestMessage := message.GuestMessage{
				Players: playerList,
				Append:  dictionary,
			}

			return message.NewSyncMessage(0, guestMessage, background)
		}
	} else {
		// 三无小号，全局战争迷雾
		playerList := generateOtherMessage(c, nil)

		guestMessage := message.GuestMessage{
			Players: playerList,
			Append:  dictionary,
		}

		return message.NewSyncMessage(0, guestMessage, background)
	}

}

// calculateReactions 计算元素反应并根据元素反应类型对伤害做出修正
func (c *Core) calculateReactions(damageCtx *context.DamageContext, targetPlayer *player) {
	for targetCharacter, damage := range damageCtx.Damage() {
		character := targetPlayer.characters[targetCharacter]
		// 为目标玩家附加伤害元素
		tempElements := c.ruleSet.ReactionCalculator.Attach(character.elements, damage.ElementType())

		// 根据目标玩家身上的元素计算反应类型
		reaction, remains := c.ruleSet.ReactionCalculator.ReactionCalculate(tempElements)
		character.elements = remains
		damageCtx.SetReaction(targetCharacter, reaction)

		// 根据元素反应类型修改伤害
		c.ruleSet.ReactionCalculator.DamageCalculate(reaction, character, damageCtx)
	}
}

// executeReactionEffect 执行元素反应的效果
func (c *Core) executeReactionEffect(reaction enum.Reaction, targetPlayer *player) {
	effect := c.ruleSet.ReactionCalculator.EffectCalculate(reaction, targetPlayer)
	c.executeCallbackModify(targetPlayer, effect)
}

// paymentCheck 检查支付费用和需求费用是否一致且玩家是否有能力支付费用
func (c *Core) paymentCheck(need, paid model.Cost, sender *player) bool {
	if !need.Equals(paid) {
		return false
	} else if !sender.holdingCost.Contains(paid) {
		return false
	} else {
		return true
	}
}

// sendSyncMessage 立即向所有玩家发送同步信息
func (c *Core) sendSyncMessage() {
	// 向所有参与战斗的玩家发送信息
	for _, playerContext := range c.room {
		syncMessage := c.generateSyncMessage(playerContext.player.uid)
		if jsonBytes, err := json.Marshal(syncMessage); err != nil {

		} else {
			playerContext.connection.Write(jsonBytes)
		}
	}

	// 向所有观战的玩家发送信息
	for viewerID, viewerConnection := range c.viewers {
		syncMessage := c.generateSyncMessage(viewerID)
		if jsonBytes, err := json.Marshal(syncMessage); err != nil {
			c.errorHandler.Handle(err)
		} else {
			viewerConnection.Write(jsonBytes)
		}
	}

	// 向所有观战的游客发送信息
	for conn, isValid := range c.guests {
		if conn != nil && isValid {
			syncMessage := c.generateSyncMessage(0)
			if jsonBytes, err := json.Marshal(syncMessage); err != nil {
				c.errorHandler.Handle(err)
			} else {
				conn.Write(jsonBytes)
			}
		}
	}
}

// filterUpdater 负责更新玩家可执行操作的协程，响应operatedChan和defeatedChan
func (c *Core) filterUpdater() {
	for existNextPlayer, actingPlayer := c.activeChain.next(); existNextPlayer; existNextPlayer, actingPlayer = c.activeChain.next() {
		if player, existPlayer := c.room[actingPlayer]; existPlayer {
			// 更新下一个玩家的状态为执行中
			c.updatePlayerStatusAndCoreFilter(player.player, enum.PlayerStatusActing)
			c.actingPlayer = actingPlayer

		block:
			// 阻塞更新协程并等待同步
			for {
				select {
				case defeated := <-c.defeatedChan:
					// 让当前执行的玩家进行等待
					waitingPlayer := c.room[c.actingPlayer].player
					c.updatePlayerStatusAndCoreFilter(waitingPlayer, enum.PlayerStatusWaiting)

					// 让被击败的玩家切换被击败的角色
					defeatedPlayer := c.room[defeated.DefeatedPlayerUID].player
					c.updatePlayerStatusAndCoreFilter(defeatedPlayer, enum.PlayerStatusCharacterDefeated)
				case switched := <-c.switchedChan:
					// 让切换完毕的玩家进入等待
					switchedPlayer := c.room[switched.SwitchedPlayerUID].player
					c.updatePlayerStatusAndCoreFilter(switchedPlayer, enum.PlayerStatusWaiting)

					// 让当前行动的玩家继续行动
					continuePlayer := c.room[c.actingPlayer].player
					c.updatePlayerStatusAndCoreFilter(continuePlayer, enum.PlayerStatusActing)
				case <-c.operatedChan:
					// 让当前行动的玩家进入等待
					completedPlayer := c.room[c.actingPlayer].player
					c.updatePlayerStatusAndCoreFilter(completedPlayer, enum.PlayerStatusWaiting)
					// 当前玩家操作完毕，退出阻塞
					break block
				case <-c.exitChan:
					c.exitChan <- struct{}{}
					return
				}
			}
		} else {
			// 下一个玩家没有被框架托管，理论上不可能，致命错误
			c.errorHandler.Handle(fmt.Errorf("error occurred while handling player %v, does not exist", actingPlayer))
			c.Close()
		}
	}
}

// networkListener 负责监听网络通信的协程，从给定的websocket连接中获取信息
func (c *Core) networkListener(conn *websocket.Connection) {
	for {
		select {
		case iStream := <-conn.Read():
			var actionMessage message.ActionMessage
			if err := json.Unmarshal(iStream, &actionMessage); err != nil {
				// 收到错误数据包，不理会
				c.errorHandler.Handle(err)
			} else {
				// 将网络消息写入消息缓冲管道
				c.readChan <- actionMessage
			}
		case <-c.exitChan:
			c.exitChan <- struct{}{}
			return
		}
	}
}

// actionExecutor 负责执行玩家操作的协程
func (c *Core) actionExecutor() {
	for {
		select {
		case msg := <-c.readChan:
			if c.messageFilter(msg) {
				// 合法信息，进行处理
				c.handleMessage(msg)
			}

		case <-c.exitChan:
			c.exitChan <- struct{}{}
			return
		}
	}
}

// executeCallback 执行某玩家的回调效果
func (c *Core) executeCallbackModify(p *player, ctx *context.CallbackContext) {
	// 执行元素转换效果
	{
		if needExecuteChangeElement, changeElementResult := ctx.ChangeElementsResult(); needExecuteChangeElement {
			// 计算需要更改的元素，将增加和删除分离
			addElement, removeElement := map[enum.ElementType]uint{}, map[enum.ElementType]uint{}
			for element, amount := range changeElementResult.Cost() {
				if amount > 0 {
					addElement[element] += uint(amount)
				} else {
					removeElement[element] += uint(-amount)
				}
			}

			// 执行增加和删除操作
			removeElementCost, addElementCost := *model.NewCostFromMap(removeElement), *model.NewCostFromMap(addElement)
			p.holdingCost.Pay(removeElementCost)
			p.holdingCost.Add(addElementCost)
		}
	}

	// 执行充能更改效果
	{
		if needChangeCharge, changeChargeResult := ctx.ChangeChargeResult(); needChangeCharge {
			// 全局修正充能效果
			p.globalChargeModifiers.Execute(changeChargeResult)

			for characterID := range changeChargeResult.Charge() {
				character := p.characters[characterID]
				// 本地修正充能效果
				character.localChargeModifiers.Execute(changeChargeResult)

				// 执行充能结果
				chargeAmount := changeChargeResult.Charge()[characterID]
				if chargeAmount > 0 {
					// 充能，需要考虑上界
					if character.currentMP+uint(chargeAmount) > character.maxMP {
						character.currentMP = character.maxMP
					} else {
						character.currentMP += uint(chargeAmount)
					}
				} else if chargeAmount < 0 {
					// 扣能量，需要考虑下界
					if chargeAmount+int(character.currentMP) < 0 {
						character.currentMP = 0
					} else {
						character.currentMP = uint(chargeAmount + int(character.currentMP))
					}
				} else {
					// 0值，不改变
					continue
				}
			}
		}
	}

	// 执行修正修改结果
	{
		if needChangeModifiers, changeModifiersResult := ctx.ChangeModifiersResult(); needChangeModifiers {
			// 修改全局修正
			{
				if changeModifiersResult.AddGlobalChargeModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalChargeModifiers() {
						p.globalChargeModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.AddGlobalCostModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalCostModifiers() {
						p.globalCostModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.AddGlobalDefenceModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalDefenceModifiers() {
						p.globalDefenceModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.AddGlobalDirectAttackModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalDirectAttackModifiers() {
						p.globalDirectAttackModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.AddGlobalFinalAttackModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalFinalAttackModifiers() {
						p.globalFinalAttackModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.AddGlobalHealModifiers() != nil {
					for _, mdf := range changeModifiersResult.AddGlobalHealModifiers() {
						p.globalHealModifiers.Append(mdf)
					}
				}

				if changeModifiersResult.RemoveGlobalChargeModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalChargeModifiers() {
						p.globalChargeModifiers.Remove(mdf.ID())
					}
				}

				if changeModifiersResult.RemoveGlobalCostModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalCostModifiers() {
						p.globalCostModifiers.Remove(mdf.ID())
					}
				}

				if changeModifiersResult.RemoveGlobalDefenceModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalDefenceModifiers() {
						p.globalDefenceModifiers.Remove(mdf.ID())
					}
				}

				if changeModifiersResult.RemoveGlobalDirectAttackModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalDirectAttackModifiers() {
						p.globalDirectAttackModifiers.Remove(mdf.ID())
					}
				}

				if changeModifiersResult.RemoveGlobalFinalAttackModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalFinalAttackModifiers() {
						p.globalFinalAttackModifiers.Remove(mdf.ID())
					}
				}

				if changeModifiersResult.RemoveGlobalHealModifiers() != nil {
					for _, mdf := range changeModifiersResult.RemoveGlobalHealModifiers() {
						p.globalHealModifiers.Remove(mdf.ID())
					}
				}
			}

			// 修改角色本地修正
			for _, character := range p.characters {
				if changeModifiersResult.AddLocalChargeModifiers() != nullChargeModifiers {
					localChargeModifiers := changeModifiersResult.AddLocalChargeModifiers().Get(character.id)
					for _, localChargeModifier := range localChargeModifiers {
						character.localChargeModifiers.Append(localChargeModifier)
					}
				}

				if changeModifiersResult.AddLocalHealModifiers() != nullHealModifiers {
					localHealModifiers := changeModifiersResult.AddLocalHealModifiers().Get(character.id)
					for _, localHealModifier := range localHealModifiers {
						character.localHealModifiers.Append(localHealModifier)
					}
				}

				if changeModifiersResult.AddLocalCostModifiers() != nullCostModifiers {
					localCostModifiers := changeModifiersResult.AddLocalCostModifiers().Get(character.id)
					for _, localCostModifier := range localCostModifiers {
						character.localCostModifiers.Append(localCostModifier)
					}
				}

				if changeModifiersResult.AddLocalDefenceModifiers() != nullDefenceModifiers {
					localDefenceModifiers := changeModifiersResult.AddLocalDefenceModifiers().Get(character.id)
					for _, localDefenceModifier := range localDefenceModifiers {
						character.localDefenceModifiers.Append(localDefenceModifier)
					}
				}

				if changeModifiersResult.AddLocalDirectAttackModifiers() != nullDirectAttackModifiers {
					localDirectAttackModifiers := changeModifiersResult.AddLocalDirectAttackModifiers().Get(character.id)
					for _, localDirectAttackModifier := range localDirectAttackModifiers {
						character.localDirectAttackModifiers.Append(localDirectAttackModifier)
					}
				}

				if changeModifiersResult.AddLocalFinalAttackModifiers() != nullFinalAttackModifiers {
					localFinalAttackModifiers := changeModifiersResult.AddLocalFinalAttackModifiers().Get(character.id)
					for _, localFinalAttackModifier := range localFinalAttackModifiers {
						character.localFinalAttackModifiers.Append(localFinalAttackModifier)
					}
				}

				if changeModifiersResult.RemoveLocalChargeModifiers() != nullChargeModifiers {
					localChargeModifiers := changeModifiersResult.RemoveLocalChargeModifiers().Get(character.id)
					for _, localChargeModifier := range localChargeModifiers {
						character.localChargeModifiers.Remove(localChargeModifier.ID())
					}
				}

				if changeModifiersResult.RemoveLocalHealModifiers() != nullHealModifiers {
					localHealModifiers := changeModifiersResult.RemoveLocalHealModifiers().Get(character.id)
					for _, localHealModifier := range localHealModifiers {
						character.localHealModifiers.Remove(localHealModifier.ID())
					}
				}

				if changeModifiersResult.RemoveLocalCostModifiers() != nullCostModifiers {
					localCostModifiers := changeModifiersResult.RemoveLocalCostModifiers().Get(character.id)
					for _, localCostModifier := range localCostModifiers {
						character.localCostModifiers.Remove(localCostModifier.ID())
					}
				}

				if changeModifiersResult.RemoveLocalDefenceModifiers() != nullDefenceModifiers {
					localDefenceModifiers := changeModifiersResult.RemoveLocalDefenceModifiers().Get(character.id)
					for _, localDefenceModifier := range localDefenceModifiers {
						character.localDefenceModifiers.Remove(localDefenceModifier.ID())
					}
				}

				if changeModifiersResult.RemoveLocalDirectAttackModifiers() != nullDirectAttackModifiers {
					localDirectAttackModifiers := changeModifiersResult.RemoveLocalDirectAttackModifiers().Get(character.id)
					for _, localDirectAttackModifier := range localDirectAttackModifiers {
						character.localDirectAttackModifiers.Remove(localDirectAttackModifier.ID())
					}
				}

				if changeModifiersResult.RemoveLocalFinalAttackModifiers() != nullFinalAttackModifiers {
					localFinalAttackModifiers := changeModifiersResult.RemoveLocalFinalAttackModifiers().Get(character.id)
					for _, localFinalAttackModifier := range localFinalAttackModifiers {
						character.localFinalAttackModifiers.Remove(localFinalAttackModifier.ID())
					}
				}
			}
		}
	}

	// 执行活动状态更改
	{
		if needChangeOperated, result := ctx.ChangeOperatedResult(); needChangeOperated {
			p.operated = result
		}
	}

	// 执行角色切换结果
	{
		if needSwitchCharacter, result := ctx.SwitchCharacterResult(); needSwitchCharacter {
			if p.characters[result].status == enum.CharacterStatusBackground {
				if p.characters[p.activeCharacter].status != enum.CharacterStatusDefeated {
					// 当前前台角色不是被击败，将其更改为后台角色
					p.characters[p.activeCharacter].status = enum.CharacterStatusBackground
				}

				// 需要切换的角色是后台角色，将其更改为前台角色
				p.characters[result].status = enum.CharacterStatusActive
				p.activeCharacter = result
			}
		}
	}

	// 执行获取卡牌效果
	{
		if needGetCard, result := ctx.GetCardsResult(); needGetCard && result > 0 {
			for i := uint(0); i < result; i++ {
				if card, success := p.cardDeck.GetOne(); success {
					// 卡牌足够的话，将卡牌加入手牌
					p.holdingCards[card.ID()] = card
				} else {
					// 卡牌不够的话，下次一定
					break
				}
			}
		}
	}

	// 执行获取特定卡牌效果
	{
		if needFindCard, target := ctx.GetFindCardResult(); needFindCard {
			if card, success := p.cardDeck.FindOne(target); success {
				// 可以获取卡牌的话，将卡牌添加进手牌
				p.holdingCards[card.ID()] = card
			}
		}
	}

	// 执行元素附着
	{
		if needAttachElement, attachment := ctx.AttachElementResult(); needAttachElement {
			for target, element := range attachment {
				targetCharacter := p.characters[target]
				tempElements := c.ruleSet.ReactionCalculator.Attach(targetCharacter.elements, element)
				reaction, remainElement := c.ruleSet.ReactionCalculator.ReactionCalculate(tempElements)
				if reaction != enum.ReactionNone {
					c.executeReactionEffect(reaction, p)
				} else {
					targetCharacter.elements = remainElement
				}
			}
		}
	}
}

// executeAttack 执行玩家的攻击指令
func (c *Core) executeAttack(action message.AttackAction) {
	senderPlayerContext, targetPlayerContext := c.room[action.Sender], c.room[action.Target]
	var senderPlayer, targetPlayer *player = nil, nil
	var attackSkill model.AttackSkill = nil

	// 执行状态校验
	{
		// 校验玩家信息
		if senderPlayerContext == nil || targetPlayerContext == nil {
			// 不存在玩家，不处理 todo add traces
			return
		} else if senderPlayer, targetPlayer = senderPlayerContext.player, targetPlayerContext.player; senderPlayer == nil || targetPlayer == nil {
			// 玩家的对战信息未被托管，不处理 todo add traces
			return
		} else if senderPlayer.characters[senderPlayer.activeCharacter].status != enum.CharacterStatusActive {
			// 发起玩家的前台角色状态无法发起攻击，不处理
			return
		}

		// 校验技能信息
		if skill, existSkill := senderPlayer.characters[senderPlayer.activeCharacter].skills[action.Skill]; !existSkill {
			// 在协同技能中查找该技能是否存在
			existCooperativeSkill := false
			for _, cooperativeSkill := range senderPlayer.cooperativeAttacks {
				if cooperativeSkill.ID() == action.Skill {
					existCooperativeSkill = true
					break
				}
			}

			// 发起玩家的前台角色不持有该技能
			if !existCooperativeSkill {
				return
			}
		} else if attack, converted := skill.(model.AttackSkill); !converted {
			// 技能无法被转化为攻击技能，不处理
			return
		} else {
			attackSkill = attack
		}
	}

	// 扣减技能费用
	{
		baseCost, paidCost := attackSkill.Cost(), *model.NewCostFromMap(action.Paid)
		if !c.paymentCheck(baseCost, paidCost, senderPlayer) {
			// 费用不合法或玩家无力承担此次费用，不处理
			return
		} else {
			// 扣费
			senderPlayer.holdingCost.Pay(paidCost)
		}
	}

	// 填充基础伤害
	baseDamage := attackSkill.BaseDamage(targetPlayer.activeCharacter, senderPlayer.activeCharacter, targetPlayer.GetBackgroundCharacters())

	// 计算伤害修正
	{
		executeCharacter := senderPlayer.characters[senderPlayer.activeCharacter]

		// 先执行攻击方直接加算逻辑
		senderPlayer.globalDirectAttackModifiers.Execute(baseDamage)
		executeCharacter.localDirectAttackModifiers.Execute(baseDamage)

		// 计算元素反应加成
		c.calculateReactions(baseDamage, targetPlayer)

		// 最后执行攻击方最终乘算逻辑
		executeCharacter.localFinalAttackModifiers.Execute(baseDamage)
		senderPlayer.globalFinalAttackModifiers.Execute(baseDamage)

		// 计算防御方减伤逻辑
		for targetCharacter, damage := range baseDamage.Damage() {
			if targetCharacter != targetPlayer.activeCharacter {
				if damage.ElementType() != enum.ElementNone {
					// 后台角色承伤，且伤害有元素效果，意味着不是穿透伤害，享受角色防御减伤
					targetPlayer.characters[targetCharacter].localDefenceModifiers.Execute(baseDamage)
				} else {
					// 穿透伤害不享受角色防御减伤
					continue
				}
			} else {
				// 前台角色享受角色防御减伤
				targetPlayer.characters[targetCharacter].localDefenceModifiers.Execute(baseDamage)
			}
		}
		targetPlayer.globalDefenceModifiers.Execute(baseDamage)
	}

	// 执行伤害结果
	{
		// 扣减生命
		for targetCharacter, damage := range baseDamage.Damage() {
			character := targetPlayer.characters[targetCharacter]
			if character.currentHP <= damage.Amount() {
				character.currentHP = 0
				character.status = enum.CharacterStatusDefeated
			} else {
				character.currentHP -= damage.Amount()
			}
		}

		// 执行元素反应效果
		c.executeReactionEffect(baseDamage.GetTargetCharacterReaction(), targetPlayer)

	}

	// 执行阻塞逻辑
	if targetPlayer.characters[targetPlayer.activeCharacter].status == enum.CharacterStatusDefeated {
		c.defeatedChan <- SyncDefeatedCharacterMessage{DefeatedPlayerUID: targetPlayer.uid}
		<-c.switchedChan
	}

	// 执行攻击、防御回调
	attackCallbackContext, defenceCallbackContext := context.NewCallbackContext(), context.NewCallbackContext()
	senderPlayer.callbackEvents.Call(enum.AfterAttack, attackCallbackContext)
	targetPlayer.callbackEvents.Call(enum.AfterDefence, defenceCallbackContext)
	c.executeCallbackModify(senderPlayer, attackCallbackContext)
	c.executeCallbackModify(targetPlayer, defenceCallbackContext)
}

// executeSwitch 执行玩家的切换角色指令
func (c *Core) executeSwitch(action message.SwitchAction) {
	switchPlayerContext, hasSwitchPlayer := c.room[action.Sender]
	paidCost := *model.NewCostFromMap(action.Paid)
	if !hasSwitchPlayer || switchPlayerContext == nil {
		// 玩家没被托管，不处理
		return
	} else if c.paymentCheck(*model.NewCostFromMap(c.ruleSet.GameOptions.SwitchCost), paidCost, switchPlayerContext.player) {
		// 玩家无法支付切换角色的费用，不处理
		return
	} else if _, hasTargetCharacter := switchPlayerContext.player.characters[action.Target]; !hasTargetCharacter {
		// 玩家没有被切换的目标角色，不处理
		return
	} else if switchPlayerContext.player.characters[action.Target].status == enum.CharacterStatusDefeated {
		// 目标角色已经被击败无法切换，不处理
		return
	}

	// 执行切换角色操作
	switchPlayer := switchPlayerContext.player
	if activeStatus := switchPlayer.characters[switchPlayer.activeCharacter].status; activeStatus != enum.CharacterStatusDefeated && activeStatus != enum.CharacterStatusDisabled {
		// 不是被击败状态和无法操作状态的角色就将其切换为后台状态，否则保持原有状态
		switchPlayer.characters[switchPlayer.activeCharacter].status = enum.CharacterStatusBackground
	}
	switchPlayer.activeCharacter = action.Target
	if targetStatus := switchPlayer.characters[switchPlayer.activeCharacter].status; targetStatus != enum.CharacterStatusDisabled {
		// 不是无法操作状态的角色就将其切换为前台状态，否则保持原有状态
		switchPlayer.characters[switchPlayer.activeCharacter].status = enum.CharacterStatusActive
	}

	// todo: callback
}

// executeBurnCard 执行玩家元素转换指令
func (c *Core) executeBurnCard(action message.BurnCardAction) {
	executePlayerContext, hasExecutePlayer := c.room[action.Sender]
	if !hasExecutePlayer || executePlayerContext == nil {
		// 玩家没被托管，不处理
		return
	} else if action.Paid < enum.ElementStartIndex || action.Paid > enum.ElementEndIndex {
		// 支付的被转化元素非七元素，不处理
		return
	}

	executePlayer := executePlayerContext.player
	if paidElementCount, hasPaidElement := executePlayer.holdingCost.Costs()[action.Paid]; !hasPaidElement || paidElementCount == 0 {
		// 玩家不持有被转化元素，不处理
		return
	} else if _, hasCard := executePlayer.holdingCards[action.Card]; !hasCard {
		// 玩家不持有支付卡牌，不处理
		return
	} else {
		// 执行元素转化
		delete(executePlayer.holdingCards, action.Card)

		paidCost := *model.NewCostFromMap(map[enum.ElementType]uint{action.Paid: 1})
		obtainedCost := *model.NewCostFromMap(map[enum.ElementType]uint{executePlayer.characters[executePlayer.activeCharacter].vision: 1})
		executePlayer.holdingCost.Pay(paidCost)
		executePlayer.holdingCost.Add(obtainedCost)
	}
}

// executeUseCard 执行玩家使用卡牌指令
func (c *Core) executeUseCard(action message.UseCardAction) {

}

// executeReRoll 执行玩家重掷骰子指令
func (c *Core) executeReRoll(action message.ReRollAction) {
	reRollPlayerContext, hasExecutePlayer := c.room[action.Sender]
	if !hasExecutePlayer || reRollPlayerContext == nil {
		// 玩家没被托管，不处理
		return
	}

	executePlayer, droppedCost := reRollPlayerContext.player, model.NewCostFromMap(action.Dropped)
	if executePlayer.holdingCost.Contains(*droppedCost) {
		// 正常请求，正常处理
		executePlayer.holdingCost.Pay(*droppedCost)
		reRollCost := model.NewRandomCost(droppedCost.Total())
		executePlayer.holdingCost.Add(*reRollCost)
	} else {
		// 不正常请求，不处理 todo: add traces
		return
	}
}

// executeSkipRound 执行玩家跳过回合指令
func (c *Core) executeSkipRound(action message.SkipRoundAction) {
	if action.Sender == c.actingPlayer && action.Sender != 0 {
		// 正常请求，正常处理
		c.operatedChan <- struct{}{}
	} else {
		// 不正常请求，不处理 todo: add traces
		return
	}
}

// executeConcede 执行玩家弃权指令
func (c *Core) executeConcede(actionMessage message.ConcedeAction) {
	concedePlayerContext, hasConcedePlayer := c.room[actionMessage.Sender]
	if !hasConcedePlayer || concedePlayerContext == nil {
		// 弃权玩家状态没有被托管，不处理
		return
	} else {
		// 将弃权玩家的状态设置为被击败
		c.updatePlayerStatusAndCoreFilter(concedePlayerContext.player, enum.PlayerStatusDefeated)
	}
}

// injectPlayers 根据初始化消息将实体注入到框架中
func (c *Core) injectPlayers(initializeMessage message.InitializeMessage) (success bool) {
	exist, ruleSetPersistence := persistence.RuleSetPersistence.QueryByID(initializeMessage.Options.RuleSet)
	if !exist {
		// 找不到规则集合，初始化失败
		return false
	}

	ruleSet := ruleSetPersistence.Ctor()().Rule
	if !ruleSet.ImplementationCheck() {
		// 规则集合含有未实现接口，初始化失败
		return false
	} else {
		ruleSet.GameOptions = &model.GameOptions{
			ReRollTimes: initializeMessage.Options.ReRollTime,
			StaticCost:  initializeMessage.Options.StaticElement,
			RollAmount:  initializeMessage.Options.ElementAmount,
			GetCards:    initializeMessage.Options.GetCards,
		}
	}

	for _, playerInfo := range initializeMessage.Players {
		if success, playerEntity := initPlayer(playerInfo, ruleSet); !success {
			return false
		} else {
			if _, exist := c.room[playerInfo.UID]; !exist {
				return false
			} else {
				c.room[playerInfo.UID].player = playerEntity
			}
		}
	}

	return true
}

// Close 关闭战斗核心的所有服务
func (c *Core) Close() {
	// 向链式反应注入中子
	c.exitChan <- struct{}{}

	// 关闭websocket连接
	for _, playerContext := range c.room {
		playerContext.connection.Close()
	}

	// 关闭各种管道
	close(c.readChan)
	close(c.defeatedChan)
	close(c.switchedChan)
}

// Serve 启动战斗核心的所有服务
func (c *Core) Serve() {
	// 启动所有websocket监听
	for _, playerContext := range c.room {
		go c.networkListener(playerContext.connection)
	}
}

// generateSelfMessage 根据core当前的状态生成一个对player发送的其自己的同步信息
func generateSelfMessage(c *Core, player *player) (selfMessage message.Self) {
	var characterList []message.Character
	for _, character := range player.characters {
		var equipmentList []message.Equipment
		for _, equipment := range character.equipments {
			// todo: finish equipment module
			equipmentList = append(equipmentList, message.Equipment{ID: equipment})
		}

		var modifierList []message.Modifier
		for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeAttack) {
			modifierList = append(modifierList, message.Modifier{
				ID:   mdf,
				Type: enum.ModifierTypeAttack,
			})
		}
		for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeCharge) {
			modifierList = append(modifierList, message.Modifier{
				ID:   mdf,
				Type: enum.ModifierTypeCharge,
			})
		}
		for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeCost) {
			modifierList = append(modifierList, message.Modifier{
				ID:   mdf,
				Type: enum.ModifierTypeCost,
			})
		}
		for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeDefence) {
			modifierList = append(modifierList, message.Modifier{
				ID:   mdf,
				Type: enum.ModifierTypeDefence,
			})
		}
		for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeHeal) {
			modifierList = append(modifierList, message.Modifier{
				ID:   mdf,
				Type: enum.ModifierTypeHeal,
			})
		}

		characterList = append(characterList, message.Character{
			ID:         character.id,
			MP:         character.currentMP,
			HP:         character.currentHP,
			Equipments: equipmentList,
			Modifiers:  modifierList,
			Status:     character.status,
		})
	}

	var campEffectList []message.Modifier
	for _, campEffect := range player.GetGlobalModifiers(enum.ModifierTypeAttack) {
		campEffectList = append(campEffectList, message.Modifier{
			ID:   campEffect,
			Type: enum.ModifierTypeAttack,
		})
	}
	for _, campEffect := range player.GetGlobalModifiers(enum.ModifierTypeCharge) {
		campEffectList = append(campEffectList, message.Modifier{
			ID:   campEffect,
			Type: enum.ModifierTypeCharge,
		})
	}
	for _, campEffect := range player.GetGlobalModifiers(enum.ModifierTypeCost) {
		campEffectList = append(campEffectList, message.Modifier{
			ID:   campEffect,
			Type: enum.ModifierTypeCost,
		})
	}
	for _, campEffect := range player.GetGlobalModifiers(enum.ModifierTypeDefence) {
		campEffectList = append(campEffectList, message.Modifier{
			ID:   campEffect,
			Type: enum.ModifierTypeDefence,
		})
	}
	for _, campEffect := range player.GetGlobalModifiers(enum.ModifierTypeHeal) {
		campEffectList = append(campEffectList, message.Modifier{
			ID:   campEffect,
			Type: enum.ModifierTypeHeal,
		})
	}

	var cardList []uint
	for id := range player.holdingCards {
		cardList = append(cardList, id)
	}

	var legalActions []enum.ActionType
	for action, isLegal := range c.room[player.uid].filter {
		if isLegal {
			legalActions = append(legalActions, action)
		}
	}

	return message.Self{
		Base: message.Base{
			UID:          player.uid,
			Characters:   characterList,
			CampEffect:   campEffectList,
			Cooperatives: nil,
			Summons:      nil,
			Supports:     nil,
			Events:       nil,
			RemainCards:  player.cardDeck.remain,
			LegalActions: legalActions,
			Status:       player.status,
		},
		Cost:  player.holdingCost.Costs(),
		Cards: cardList,
	}
}

// generateOtherMessage 根据core当前的状态生成一个除player外发送给其他人的同步信息
func generateOtherMessage(c *Core, player *player) (othersMessage []message.Other) {
	othersMessage = []message.Other{}

	for _, playerContext := range c.room {
		otherPlayer := playerContext.player
		if otherPlayer == player {
			continue
		}

		var characterList []message.Character
		for _, character := range otherPlayer.characters {
			// todo: finish equipment module
			var equipmentList []message.Equipment
			for _, equipment := range character.equipments {
				equipmentList = append(equipmentList, message.Equipment{
					ID:   equipment,
					Type: 0,
				})
			}

			var modifierList []message.Modifier
			for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeAttack) {
				modifierList = append(modifierList, message.Modifier{
					ID:   mdf,
					Type: enum.ModifierTypeAttack,
				})
			}
			for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeCharge) {
				modifierList = append(modifierList, message.Modifier{
					ID:   mdf,
					Type: enum.ModifierTypeCharge,
				})
			}
			for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeCost) {
				modifierList = append(modifierList, message.Modifier{
					ID:   mdf,
					Type: enum.ModifierTypeCost,
				})
			}
			for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeDefence) {
				modifierList = append(modifierList, message.Modifier{
					ID:   mdf,
					Type: enum.ModifierTypeDefence,
				})
			}
			for _, mdf := range character.GetLocalModifiers(enum.ModifierTypeHeal) {
				modifierList = append(modifierList, message.Modifier{
					ID:   mdf,
					Type: enum.ModifierTypeHeal,
				})
			}

			characterList = append(characterList, message.Character{
				ID:         character.id,
				MP:         character.currentMP,
				HP:         character.currentHP,
				Equipments: equipmentList,
				Modifiers:  modifierList,
				Status:     character.status,
			})
		}

		var campEffectList []message.Modifier
		for _, campEffect := range otherPlayer.GetGlobalModifiers(enum.ModifierTypeAttack) {
			campEffectList = append(campEffectList, message.Modifier{
				ID:   campEffect,
				Type: enum.ModifierTypeAttack,
			})
		}
		for _, campEffect := range otherPlayer.GetGlobalModifiers(enum.ModifierTypeCharge) {
			campEffectList = append(campEffectList, message.Modifier{
				ID:   campEffect,
				Type: enum.ModifierTypeCharge,
			})
		}
		for _, campEffect := range otherPlayer.GetGlobalModifiers(enum.ModifierTypeCost) {
			campEffectList = append(campEffectList, message.Modifier{
				ID:   campEffect,
				Type: enum.ModifierTypeCost,
			})
		}
		for _, campEffect := range otherPlayer.GetGlobalModifiers(enum.ModifierTypeDefence) {
			campEffectList = append(campEffectList, message.Modifier{
				ID:   campEffect,
				Type: enum.ModifierTypeDefence,
			})
		}
		for _, campEffect := range otherPlayer.GetGlobalModifiers(enum.ModifierTypeHeal) {
			campEffectList = append(campEffectList, message.Modifier{
				ID:   campEffect,
				Type: enum.ModifierTypeHeal,
			})
		}

		var legalActions []enum.ActionType
		for action, isLegal := range c.room[otherPlayer.uid].filter {
			if isLegal {
				legalActions = append(legalActions, action)
			}
		}

		other := message.Other{
			Base: message.Base{
				UID:          otherPlayer.uid,
				Characters:   characterList,
				CampEffect:   campEffectList,
				Cooperatives: nil,
				Summons:      nil,
				Supports:     nil,
				Events:       nil,
				RemainCards:  otherPlayer.cardDeck.remain,
				LegalActions: legalActions,
				Status:       otherPlayer.status,
			},
			Cost:  playerContext.player.holdingCost.Total(),
			Cards: uint(len(player.holdingCards)),
		}

		othersMessage = append(othersMessage, other)
	}

	return othersMessage
}

// generateBackgroundMessage 根据core当前的状态生成游戏当前的状态信息
func generateBackgroundMessage(c *Core) (gameMessage message.Game) {
	return message.Game{
		ActingPlayer: c.actingPlayer,
		RoundStage:   c.roundStage,
		RoundCount:   c.roundCount,
	}
}

// generateDictionary 根据core当前的状态生成EntityID-TypeID的字典
func generateDictionary(c *Core) (dictionary []message.DictionaryPair) {
	dictionary = []message.DictionaryPair{}
	for id, typeID := range c.entities {
		dictionary = append(dictionary, message.DictionaryPair{
			TypeID:   typeID,
			EntityID: id,
		})
	}
	return dictionary
}

// initCharacter 新建并初始化一个character实体
func initCharacter(characterID, ownerID uint) (success bool, result *character) {
	exist, characterPersistence := persistence.CharacterPersistence.QueryByID(characterID)
	if !exist {
		// 找不到角色实现，初始化失败
		return false, nil
	}

	characterInfo := characterPersistence.Ctor()()
	characterSkill := map[uint]model.Skill{}
	for _, skillID := range characterInfo.Skills {
		if existSkill, skillPersistence := persistence.SkillPersistence.QueryByID(skillID); !existSkill {
			// 找不到技能实现，初始化失败
			return false, nil
		} else {
			skill := skillPersistence.Ctor()().Skill
			characterSkill[skillID] = skill
		}
	}

	character := &character{
		id:                         characterInfo.ID,
		player:                     ownerID,
		affiliation:                characterInfo.Affiliation,
		vision:                     characterInfo.Vision,
		weapon:                     characterInfo.Weapon,
		skills:                     characterSkill,
		maxHP:                      characterInfo.MaxHP,
		currentHP:                  characterInfo.MaxHP,
		maxMP:                      characterInfo.MaxMP,
		currentMP:                  0,
		status:                     enum.CharacterStatusReady,
		elements:                   []enum.ElementType{},
		satiety:                    false,
		equipments:                 map[enum.EquipmentType]uint{},
		localDirectAttackModifiers: modifier.NewChain[context.DamageContext](),
		localFinalAttackModifiers:  modifier.NewChain[context.DamageContext](),
		localDefenceModifiers:      modifier.NewChain[context.DamageContext](),
		localChargeModifiers:       modifier.NewChain[context.ChargeContext](),
		localHealModifiers:         modifier.NewChain[context.HealContext](),
		localCostModifiers:         modifier.NewChain[context.CostContext](),
	}

	return true, character
}

// initPlayer 新建并初始化一个player实体
func initPlayer(matchingMessage message.MatchingMessage, ruleSet model.RuleSet) (success bool, result *player) {
	if existPlayer, _ := persistence.PlayerPersistence.QueryByID(matchingMessage.UID); !existPlayer {
		// 不存在玩家信息，初始化失败
		return false, nil
	}

	var characterList []uint
	characterMap := map[uint]*character{}
	for _, characterID := range matchingMessage.Characters {
		if initCharacterSuccess, character := initCharacter(characterID, matchingMessage.UID); !initCharacterSuccess {
			// 初始化角色失败
			return false, nil
		} else {
			characterList = append(characterList, characterID)
			characterMap[characterID] = character
		}
	}

	var cardList []model.Card
	for _, cardID := range matchingMessage.CardDeck {
		if existCard, cardPersistence := persistence.CardPersistence.QueryByID(cardID); !existCard {
			// 不存在卡牌，初始化失败
			return false, nil
		} else {
			cardList = append(cardList, cardPersistence.Ctor()().Card)
		}
	}

	player := &player{
		uid:                         matchingMessage.UID,
		status:                      enum.PlayerStatusInitialized,
		operated:                    false,
		reRollTimes:                 ruleSet.GameOptions.ReRollTimes,
		staticCost:                  model.NewCostFromMap(ruleSet.GameOptions.StaticCost),
		holdingCost:                 model.NewCost(),
		cardDeck:                    NewCardDeck(cardList),
		holdingCards:                map[uint]model.Card{},
		activeCharacter:             0,
		characters:                  characterMap,
		characterList:               characterList,
		summons:                     map[uint]Summon{},
		summonList:                  []uint{},
		supports:                    map[uint]Support{},
		supportList:                 []uint{},
		globalDirectAttackModifiers: modifier.NewChain[context.DamageContext](),
		globalFinalAttackModifiers:  modifier.NewChain[context.DamageContext](),
		globalDefenceModifiers:      modifier.NewChain[context.DamageContext](),
		globalChargeModifiers:       modifier.NewChain[context.ChargeContext](),
		globalHealModifiers:         modifier.NewChain[context.HealContext](),
		globalCostModifiers:         modifier.NewChain[context.CostContext](),
		cooperativeAttacks:          map[enum.TriggerType]model.CooperativeSkill{},
		callbackEvents:              event.NewEventMap(),
	}

	return true, player
}
