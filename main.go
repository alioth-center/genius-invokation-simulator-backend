package main

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

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
	return 10
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

func main() {
	ruleSet := entity.NewEmptyRuleSet()

	ganyu1 := entity.NewCharacter(1, Ganyu{id: 2333}, ruleSet)
	fmt.Printf("initializing characters: %+v\n", ganyu1)

	ganyu2 := entity.NewCharacter(2, Ganyu{id: 3333}, ruleSet)
	fmt.Printf("initializing characters: %+v\n", ganyu2)

	modifiers := &context.ModifierContext{}
	modifiers.AddLocalDefenceModifier(3333, DefenceBuff{id: 2233})
	fmt.Printf("add defence buff: %+v\n", modifiers)
	ganyu2.ExecuteModify(modifiers)

	attack1 := ganyu1.ExecuteAttack(1, ganyu2.ID(), []uint{})
	fmt.Printf("initializing attack context: %+v\n", attack1)

	ganyu2.ExecuteDefence(attack1)
	fmt.Printf("be attacked: %+v\n", ganyu2)

	attack2 := ganyu1.ExecuteAttack(1, ganyu2.ID(), []uint{})
	fmt.Printf("initializing attack context: %+v\n", attack2)

	ganyu2.ExecuteDefence(attack2)
	fmt.Printf("be attacked: %+v\n", ganyu2)
}
