package main

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type EmptyRule struct{}

func (e EmptyRule) ReactionCalculate(types []enum.ElementType) (reaction enum.Reaction, elementRemains []enum.ElementType) {
	return enum.ReactionNone, types
}

func (e EmptyRule) DamageCalculate(reaction enum.Reaction, targetCharacter uint, ctx *context.DamageContext) {
}

func (e EmptyRule) EffectCalculate(reaction enum.Reaction, targetPlayer entity.Player) (ctx *context.CallbackContext) {
	return nil
}

func (e EmptyRule) Attach(originalElements []enum.ElementType, newElement enum.ElementType) (resultElements []enum.ElementType) {
	return originalElements
}

func (e EmptyRule) Relative(reaction enum.Reaction, relativeElement enum.ElementType) bool {
	return false
}

type Ganyu struct {
	id uint
}

func (e Ganyu) ID() uint {
	return e.id
}

func (e Ganyu) Affiliation() enum.Affiliation {
	return enum.AffiliationLiyue
}

func (e Ganyu) Vision() enum.ElementType {
	return enum.ElementCryo
}

func (e Ganyu) Weapon() enum.WeaponType {
	return enum.WeaponBow
}

func (e Ganyu) MaxHP() uint {
	return 16
}

func (e Ganyu) MaxMP() uint {
	return 2
}

func (e Ganyu) Skills() map[uint]entity.Skill {
	return map[uint]entity.Skill{1: Shuanghuashi{id: 1}}
}

type Shuanghuashi struct {
	id uint
}

func (e Shuanghuashi) ID() uint {
	return e.id
}

func (e Shuanghuashi) Type() enum.SkillType {
	return enum.SkillNormalAttack
}

func (e Shuanghuashi) Cost() entity.Cost {
	return *entity.NewCost()
}

func (e Shuanghuashi) BaseDamage(target, self uint, background []uint) *context.DamageContext {
	return context.NewDamageContext(e.id, self, target, background, enum.ElementNone, 10)
}

type DefenceBuff struct {
	id uint
}

func (d DefenceBuff) ID() uint {
	return d.id
}

func (d DefenceBuff) Handler() func(ctx *modifier.Context[context.DamageContext]) {
	return func(ctx *modifier.Context[context.DamageContext]) {
		ctx.Data().SubActiveDamage(4)
	}
}

func (d DefenceBuff) Clone() modifier.Modifier[context.DamageContext] {
	return d
}

func (d DefenceBuff) RoundReset() {

}

func (d DefenceBuff) Effective() bool {
	return true
}

func (d DefenceBuff) EffectLeft() uint {
	return 0
}

type playerInfo struct {
	uid        uint
	name       string
	cards      []entity.Card
	characters map[uint]entity.Character
}

func (p playerInfo) UID() uint { return p.uid }

func (p playerInfo) Name() string { return p.name }

func (p playerInfo) Cards() []entity.Card { return p.cards }

func (p playerInfo) Characters() map[uint]entity.Character { return p.characters }

func main() {
	ruleSet := entity.NewRuleSet(EmptyRule{})
	ganyu1 := entity.NewCharacter(1, Ganyu{id: 2333}, ruleSet)
	ganyu2 := entity.NewCharacter(2, Ganyu{id: 3333}, ruleSet)
	player1 := entity.NewPlayer(playerInfo{
		uid:        1,
		name:       "player1",
		cards:      []entity.Card{},
		characters: map[uint]entity.Character{2333: ganyu1},
	})
	player2 := entity.NewPlayer(playerInfo{
		uid:        2,
		name:       "player2",
		cards:      []entity.Card{},
		characters: map[uint]entity.Character{3333: ganyu2},
	})

	core := entity.NewCore(ruleSet, []entity.Player{player1, player2})
	fmt.Printf("player2.character: HP: %v Status: %v\n", ganyu2.HP(), ganyu2.Status())

	core.ExecuteAttack(player1.UID(), player2.UID(), 1)
	fmt.Printf("player2.character: HP: %v Status: %v\n", ganyu2.HP(), ganyu2.Status())
	core.ExecuteAttack(player1.UID(), player2.UID(), 1)
	fmt.Printf("player2.character: HP: %v Status: %v\n", ganyu2.HP(), ganyu2.Status())
}
