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

func newPlayerChain() *playerChain {
	return &playerChain{
		canOperated: kv.NewSimpleMap[bool](),
		queue:       []uint{},
		offset:      0,
	}
}

type Framework interface {
}

type framework struct {
	ruleSet     RuleSet
	players     kv.Map[uint, Player]
	activeChain *playerChain
	nextChain   *playerChain
}

func (f *framework) ExecuteAttack(sender uint, target uint, skill uint) {
	if f.players.Exists(sender) && f.players.Exists(target) {
		senderPlayer, targetPlayer := f.players.Get(sender), f.players.Get(target)

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
			activeReaction := enum.ReactionNone
			for targetCharacterID := range ctx.Damage() {
				_, c := targetPlayer.GetCharacter(targetCharacterID)
				reaction := c.ExecuteElementReaction()
				if targetCharacterID == targetCharacter.ID() {
					activeReaction = reaction
				}
				f.ruleSet.ReactionCalculator().DamageCalculate(reaction, targetCharacterID, ctx)
			}
			senderPlayer.ExecuteFinalAttackModifiers(ctx)
			targetPlayer.ExecuteDefence(ctx)
			if event := f.ruleSet.ReactionCalculator().EffectCalculate(activeReaction); event != nil {
				targetPlayer.ExecuteCallbackModify(event)
			}

			// 执行回调流程
			senderPlayer.ExecuteAfterAttackCallback()
			targetPlayer.ExecuteAfterDefenceCallback()
		}
	}
}
