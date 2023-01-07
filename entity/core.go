package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type playerChain struct {
	canOperated kv.Map[uint, bool]
	queue       []uint
	offset      int
}

func (pc *playerChain) next() (has bool, player uint) {
	for i := pc.offset; i < len(pc.queue); i++ {
		if pc.canOperated.Get(pc.queue[i]) {
			pc.offset = i + 1
			return true, pc.queue[i]
		}
	}
	for i := 0; i < pc.offset; i++ {
		if pc.canOperated.Get(pc.queue[i]) {
			pc.offset = i + 1
			return true, pc.queue[i]
		}
	}

	return false, 0
}

func (pc *playerChain) complete(player uint) {
	pc.canOperated.Set(player, false)
}

func (pc *playerChain) empty() {
	pc.canOperated = kv.NewSimpleMap[bool]()
	pc.queue = []uint{}
	pc.offset = 0
}

func (pc *playerChain) add(player Player) {
	pc.canOperated.Set(player.UID(), true)
	pc.queue = append(pc.queue, player.UID())
}

func (pc *playerChain) allActive() []uint {
	result := make([]uint, 0)
	for _, id := range pc.queue {
		if pc.canOperated.Get(id) {
			result = append(result, id)
		}
	}

	return result
}

func newPlayerChain() *playerChain {
	return &playerChain{
		canOperated: kv.NewSimpleMap[bool](),
		queue:       []uint{},
		offset:      0,
	}
}

type Core struct {
	Players     kv.Map[uint, Player]
	RoundCount  uint
	ruleSet     RuleSet
	activeChain *playerChain
	nextChain   *playerChain
}

// RoundEnd 回合结束时的结算逻辑
func (c *Core) RoundEnd() {
	// 将行动队列更新为下回合队列
	c.activeChain, c.nextChain = c.nextChain, c.activeChain
	c.nextChain.empty()
	newCallList := c.activeChain.allActive()

	// 按照行动队列执行召唤物结算
	for _, player := range newCallList {
		c.Players.Get(player).ExecuteSummonSkills()
	}

	// 按照行动队列执行结束回合结算
	for _, player := range newCallList {
		c.Players.Get(player).ExecuteRoundEndCallback()
	}

	// 回合计数器增加
	c.RoundCount++
}

// RoundStart 回合开始时的结算逻辑
func (c *Core) RoundStart() {
	newCallList := c.activeChain.allActive()

	// 按照行动队列重置所有玩家的状态
	for _, player := range newCallList {
		c.Players.Get(player).ExecuteResetCallback()
	}

	// 按照行动队列执行玩家的开始阶段
	for _, player := range newCallList {
		c.Players.Get(player).ExecuteRoundStartCallback()
	}
}

// RoundRoll 投掷阶段的计算逻辑
func (c *Core) RoundRoll() {
	newCallList := c.activeChain.allActive()
	for _, player := range newCallList {
		targetPlayer := c.Players.Get(player)
		holdingCost := targetPlayer.StaticCost()

		// 计算需要多少随机元素骰子
		remainRandomCount := int(c.ruleSet.GameOptions.RollAmount) - int(holdingCost.total)
		if remainRandomCount >= 0 {
			// 使用随机元素骰子补足缺口
			randomCost := NewRandomCost(uint(remainRandomCount))
			holdingCost.Add(*randomCost)

			targetPlayer.SetHoldingCost(holdingCost)

		} else {
			// 如果固定骰子多余可获得骰子，舍弃多出的固定元素骰子
			have, finalCount := uint(0), c.ruleSet.GameOptions.RollAmount
			finalCost := NewCost()
			for element := enum.ElementType(0); element <= enum.ElementEndIndex; element++ {
				if holdingCost.costs[element]+have > finalCount {
					finalCost.add(element, finalCount-have)
					break
				} else {
					finalCost.add(element, holdingCost.costs[element])
					have += holdingCost.costs[element]
				}
			}

			targetPlayer.SetHoldingCost(*finalCost)
		}
	}
}

func (c *Core) ExecuteReRoll(sender uint, drop map[enum.ElementType]uint) {
	if c.Players.Exists(sender) {
		c.Players.Get(sender).ExecuteElementReRoll(*NewCostFromMap(drop))
	}
}

func (c *Core) ExecutePayment(sender uint, need, paid map[enum.ElementType]uint) {
	if c.Players.Exists(sender) {
		basicCost, paidCost := NewCostFromMap(need), NewCostFromMap(paid)
		c.Players.Get(sender).ExecuteElementPayment(*basicCost, *paidCost)
	}
}

func (c *Core) ExecuteAttack(sender uint, target uint, skill uint) {
	if c.Players.Exists(sender) && c.Players.Exists(target) {
		senderPlayer, targetPlayer := c.Players.Get(sender), c.Players.Get(target)

		if has, character := senderPlayer.GetActiveCharacter(); has && character.HasSkill(skill) {
			// 填充DamageContext
			_, targetCharacter := targetPlayer.GetActiveCharacter()
			backgroundCharacters := targetPlayer.GetBackgroundCharacters()
			background := make([]uint, len(backgroundCharacters))
			for i, backgroundCharacter := range backgroundCharacters {
				background[i] = backgroundCharacter.ID()
			}
			ctx := senderPlayer.ExecuteAttack(skill, targetCharacter.ID(), background)

			// 执行攻击流程
			senderPlayer.ExecuteDirectAttackModifiers(ctx)
			targetPlayer.ExecuteElementAttachment(ctx)
			for targetCharacterID := range ctx.Damage() {
				_, executeCharacter := targetPlayer.GetCharacter(targetCharacterID)
				reaction := executeCharacter.ExecuteElementReaction()
				ctx.SetReaction(targetCharacterID, reaction)
				c.ruleSet.ReactionCalculator.DamageCalculate(reaction, targetCharacterID, ctx)
			}
			senderPlayer.ExecuteFinalAttackModifiers(ctx)
			targetPlayer.ExecuteDefence(ctx)
			if event := c.ruleSet.ReactionCalculator.EffectCalculate(ctx.GetTargetCharacterReaction(), targetPlayer); event != nil {
				targetPlayer.ExecuteCallbackModify(event)
			}

			// 执行回调流程
			senderPlayer.ExecuteAfterAttackCallback()
			targetPlayer.ExecuteAfterDefenceCallback()
		}
	}
}

func (c *Core) ExecuteBurnCard(sender uint, card uint, exchangeElement enum.ElementType) {
	if c.Players.Exists(sender) {
		c.Players.Get(sender).ExecuteBurnCard(card, exchangeElement)
	}
}

func (c *Core) ExecuteSkipRound(sender uint) {
	if c.Players.Exists(sender) {
		c.Players.Get(sender).ExecuteSkipRound()
	}
}

func (c *Core) ExecuteConcede(sender uint) {
	if c.Players.Exists(sender) {
		c.Players.Get(sender).ExecuteConcede()
	}
}

func (c *Core) ExecuteSwitchCharacter(sender uint, targetCharacter uint) {
	if c.Players.Exists(sender) {
		player := c.Players.Get(sender)
		if has, _ := player.GetCharacter(targetCharacter); has {
			player.SwitchCharacter(targetCharacter)
		}
	}
}

func (c *Core) ExecuteUseCard(sender uint, card uint) {
	if c.Players.Exists(sender) {
		player := c.Players.Get(sender)
		if player.HeldCard(card) {
			// todo: implement use card logic
			panic("not implemented yet")
		}
	}
}

func NewCore(rule RuleSet, players []Player) *Core {
	core := &Core{
		RoundCount:  0,
		ruleSet:     rule,
		Players:     kv.NewSimpleMap[Player](),
		activeChain: newPlayerChain(),
		nextChain:   newPlayerChain(),
	}

	for _, player := range players {
		core.activeChain.add(player)
		core.Players.Set(player.UID(), player)
	}

	return core
}
