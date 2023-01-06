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

	// 查询并设置规则集合
	var ruleSet entity.RuleSet
	if has, rule := db.RuleSetPersistence.QueryByID(initialize.Options.RuleSet); has {
		ruleSet = rule.Ctor()()
		ruleSet.SetOptions(entity.GameOptions{
			ReRollTimes: initialize.Options.ReRollTime,
			StaticCost:  initialize.Options.StaticElement,
			RollAmount:  initialize.Options.ElementAmount,
			GetCards:    initialize.Options.GetCards,
		})
	} else {
		return false, nil
	}

	// 查询并填充玩家信息
	playerList := make([]entity.Player, len(initialize.Players))
	for i, player := range initialize.Players {
		if has, playerRecord := db.PlayerPersistence.QueryByID(player.UID); has {
			playerInfo := playerRecord.Ctor()()

			// 注入玩家的角色信息
			characterList := make([]entity.Character, len(player.Characters))
			for ii, characterID := range player.Characters {
				if ok, character := db.CharacterPersistence.QueryByID(characterID); ok {
					characterInfo := character.Ctor()()

					// 注入角色的技能信息
					skillList := make([]entity.Skill, len(character.Skills))
					for jj, skillID := range character.Skills {
						if okk, skill := db.SkillPersistence.QueryByID(skillID); okk {
							skillEntity := skill.Ctor()()
							skillList[jj] = skillEntity
						}
					}

					// 实例化角色列表
					characterInfo.SetSkills(skillList)
					characterEntity := entity.NewCharacter(player.UID, characterInfo, ruleSet)
					characterList[ii] = characterEntity
				} else {
					return false, nil
				}
			}

			// 注入玩家的卡牌信息
			cardList := make([]entity.Card, len(player.CardDeck))
			for ii, cardID := range player.CardDeck {
				if ok, card := db.CardPersistence.QueryByID(cardID); ok {
					cardEntity := card.Ctor()()
					cardList[ii] = cardEntity
				} else {
					return false, nil
				}
			}

			// 注入玩家的基础信息
			playerInfo.SetUID(player.UID)
			playerInfo.SetCharacters(characterList)
			playerInfo.SetCards(cardList)

			playerList[i] = entity.NewPlayer(playerInfo)
		} else {
			return false, nil
		}
	}

	// 生成战斗核心
	framework.core = entity.NewCore(ruleSet, playerList)
	framework.update()

	return true, framework
}
