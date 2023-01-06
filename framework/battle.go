package framework

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/message"
	db "github.com/sunist-c/genius-invokation-simulator-backend/persistence"
)

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
	}
)

func newFilter() filter {
	return cachedFilter[enum.PlayerStatusViewing]
}

type Battle struct {
	core   *entity.Core
	filter kv.Map[uint, filter]
	in     chan message.ActionMessage
	out    chan message.SyncMessage
	exit   chan struct{}
}

func (b *Battle) update() {
	b.core.Players.Range(func(id uint, player entity.Player) bool {
		b.filter.Set(id, cachedFilter[player.Status()])
		return true
	})
}

func (b *Battle) limit(msg message.ActionMessage) bool {
	if b.filter.Exists(msg.Sender) {
		has, result := b.filter.Get(msg.Sender)[msg.Type]
		return has && result
	} else {
		return false
	}
}

func (b *Battle) serve() {
	select {
	case msg := <-b.in:
		if !b.limit(msg) {
			// todo: handle player message
		}
	case <-b.exit:
		defer close(b.in)
		defer close(b.out)
		return
	}
}

func NewBattle(initialize message.InitializeMessage) (success bool, framework *Battle) {
	// 初始化战斗框架
	framework = &Battle{
		filter: kv.NewSyncMap[filter](),
	}

	// todo: implement persistence module
	//// 查询并填充玩家信息
	playerList := make([]entity.Player, len(initialize.Players))
	//for i, player := range initialize.Players {
	//	if has, info := db.PlayerPersistence.QueryByID(player.UID); has {
	//		var playerInfo entity.PlayerInfo
	//
	//		for ii, characterID := range player.Characters {
	//			if ok, character := db.CharacterPersistence.QueryByID(characterID); ok {
	//
	//			} else {
	//				return false, nil
	//			}
	//		}
	//
	//		for ii, cardID := range player.CardDeck {
	//			if ok, card := db.CardPersistence.QueryByID(cardID); ok {
	//
	//			} else {
	//				return false, nil
	//			}
	//		}
	//
	//		playerList[i] = entity.NewPlayer(playerInfo)
	//	} else {
	//		return false, nil
	//	}
	//}

	// 查询并设置规则集合
	var ruleSet entity.RuleSet
	if has, rule := db.RuleSetPersistence.QueryByID(initialize.Options.RuleSet); has {
		ruleSet = entity.NewRuleSet(rule.ReactionCalculator, entity.GameOptions{
			ReRollTimes: initialize.Options.ReRollTime,
			StaticCost:  initialize.Options.StaticElement,
			RollAmount:  initialize.Options.ElementAmount,
			GetCards:    initialize.Options.GetCards,
		})
	} else {
		return false, nil
	}

	// 注入战斗核心
	framework.core = entity.NewCore(ruleSet, playerList)
	framework.update()

	return true, framework
}
