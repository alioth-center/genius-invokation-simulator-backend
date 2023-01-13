package entity

import (
	"encoding/json"
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/event"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/websocket"
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

type playerContext struct {
	player     *player
	connection *websocket.Connection
	filter     filter
}

type Core struct {
	room         map[uint]*playerContext           // room 房间信息，包括玩家、合法操作
	viewers      map[uint]*websocket.Connection    // viewers 观战者的连接集合
	guests       []*websocket.Connection           // guests 匿名观战者的连接集合
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
	errChan      chan error                        // errChan 错误管道，发生错误时向此管道写入信息，可能被其他结构共用
}

// updatePlayerStatusAndCoreFilter 更新玩家状态与玩家可操作列表，没有做校验，请在调用前校验player不为nil且在filter中有记录
func (c *Core) updatePlayerStatusAndCoreFilter(player *player, status enum.PlayerStatus) {
	player.status = status
	c.room[player.uid].filter = cachedFilter[status]
}

// messageFilter 过滤不合法的操作类型，若玩家发送的操作信息合法，则legal为真
func (c *Core) messageFilter(msg message.ActionMessage) (legal bool) {
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
		c.errChan <- fmt.Errorf("unknown action %d, sent by %v", msg.Type, msg.Sender)
	}

	// 处理完毕，需要向所有玩家同步信息
	c.sendSyncMessage()
}

// generateSyncMessage 生成某玩家收到的同步信息
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

// sendSyncMessage 立即向所有玩家发送同步信息
func (c *Core) sendSyncMessage() {
	// 向所有参与战斗的玩家发送信息
	for _, playerContext := range c.room {
		syncMessage := c.generateSyncMessage(playerContext.player.uid)
		if jsonBytes, err := json.Marshal(syncMessage); err != nil {
			c.errChan <- err
		} else {
			playerContext.connection.Write(jsonBytes)
		}
	}

	// 向所有观战的玩家发送信息
	for viewerID, viewerConnection := range c.viewers {
		syncMessage := c.generateSyncMessage(viewerID)
		if jsonBytes, err := json.Marshal(syncMessage); err != nil {
			c.errChan <- err
		} else {
			viewerConnection.Write(jsonBytes)
		}
	}

	// 向所有观战的游客发送信息
	for _, conn := range c.guests {
		syncMessage := c.generateSyncMessage(0)
		if jsonBytes, err := json.Marshal(syncMessage); err != nil {
			c.errChan <- err
		} else {
			conn.Write(jsonBytes)
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
			c.errChan <- fmt.Errorf("error occurred while handling player %v, does not exist", actingPlayer)
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
				c.errChan <- err
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

// executeAttack 执行攻击行动
func (c *Core) executeAttack(action message.AttackAction) {

}

func (c *Core) executeSwitch(action message.SwitchAction) {

}

func (c *Core) executeBurnCard(action message.BurnCardAction) {

}

func (c *Core) executeUseCard(action message.UseCardAction) {

}

func (c *Core) executeReRoll(action message.ReRollAction) {
	executePlayer := c.room[action.Sender].player
	if executePlayer.holdingCost.Contains(*model.NewCostFromMap(action.Dropped)) {

	}
}

func (c *Core) executeSkipRound(action message.SkipRoundAction) {

}

func (c *Core) executeConcede(actionMessage message.ConcedeAction) {

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

func generateBackgroundMessage(c *Core) (gameMessage message.Game) {
	return message.Game{
		ActingPlayer: c.actingPlayer,
		RoundStage:   c.roundStage,
		RoundCount:   c.roundCount,
	}
}

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

func initCharacter(characterID, ownerID uint, ruleSet model.RuleSet) (success bool, result *character) {
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
		ruleSet:                    ruleSet,
	}

	return true, character
}

func initPlayer(matchingMessage message.MatchingMessage, ruleSet model.RuleSet) (success bool, result *player) {
	if existPlayer, _ := persistence.PlayerPersistence.QueryByID(matchingMessage.UID); !existPlayer {
		// 不存在玩家信息，初始化失败
		return false, nil
	}

	var characterList []uint
	var characterMap map[uint]*character
	for _, characterID := range matchingMessage.Characters {
		if initCharacterSuccess, character := initCharacter(characterID, matchingMessage.UID, ruleSet); !initCharacterSuccess {
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
		cooperativeAttacks:          []model.CooperativeSkill{},
		callbackEvents:              event.NewEventMap(),
	}

	return true, player
}
