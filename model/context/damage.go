package context

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type Damage struct {
	elementType enum.ElementType
	amount      uint
}

func (d *Damage) add(amount uint) {
	d.amount += amount
}

func (d *Damage) sub(amount uint) {
	if d.amount > amount {
		d.amount -= amount
	} else {
		d.amount = 0
	}
}

func (d *Damage) change(element enum.ElementType) {
	d.elementType = element
}

// Amount 伤害的数值，只读
func (d Damage) Amount() uint {
	return d.amount
}

// ElementType 伤害的元素类型，只读
func (d Damage) ElementType() enum.ElementType {
	return d.elementType
}

type DamageContext struct {
	skillID              uint
	sendPlayer           uint
	targetCharacter      uint
	backgroundCharacters []uint
	damages              map[uint]*Damage
}

// AddActiveDamage 增加对目标玩家前台角色的伤害数值
func (d *DamageContext) AddActiveDamage(amount uint) {
	d.damages[d.targetCharacter].add(amount)
}

// AddPenetratedDamage 增加对目标玩家所有后台角色的穿透伤害数值
func (d *DamageContext) AddPenetratedDamage(amount uint) {
	for _, character := range d.backgroundCharacters {
		d.damages[character].add(amount)
	}
}

// SubActiveDamage 降低对目标玩家前台角色的伤害数值
func (d *DamageContext) SubActiveDamage(amount uint) {
	d.damages[d.targetCharacter].sub(amount)
}

// SubPenetratedDamage 降低对目标玩家所有后台角色的穿透伤害数值
func (d *DamageContext) SubPenetratedDamage(amount uint) {
	for _, character := range d.backgroundCharacters {
		d.damages[character].sub(amount)
	}
}

// ChangeElementType 修改对目标玩家前台角色的伤害元素类型
func (d *DamageContext) ChangeElementType(element enum.ElementType) {
	d.damages[d.targetCharacter].change(element)
}

// Damage 返回DamageContext携带的伤害信息，只读
func (d *DamageContext) Damage() map[uint]Damage {
	result := map[uint]Damage{}
	for target, damage := range d.damages {
		result[target] = *damage
	}

	return result
}

type DamageModifier func(ctx *modifier.Context[DamageContext])
