/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "character.go" LastUpdatedAt 2022/12/12 17:41:12
 */

package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
)

type Character struct {
	player             *Player
	ID                 uint
	Name               string
	Title              string
	Description        string
	Affiliation        definition.Affiliation
	VisionType         definition.Element
	WeaponType         definition.Weapon
	Skills             []ISkill
	MaxHealthPoint     uint
	CurrentHealthPoint uint
	MaxMagicPoint      uint
	CurrentMagicPoint  uint
	ElementsAttached   []definition.Element
	Equipments         map[definition.EquipmentType]ICard
	Status             definition.CharacterStatus
	Satiety            bool
	CostModifiers      *ModifierChain[CostContext]
	AttackModifiers    *ModifierChain[AttackDamageContext]
	CallbackChain      map[definition.Trigger]*ModifierChain[CharacterContext]
}

func (c *Character) DefenceModifiers() *ModifierChain[DefenceDamageContext] {
	return c.player.DefenceModifiers
}

// Attack 角色释放攻击，生成相应的伤害数据
func (c *Character) Attack(target *Player, skill IAttackSkill) *AttackDamageContext {
	damageData := skill.Attack(target)
	ctx := NewContext(c.AttackModifiers, damageData)
	ctx.Continue()
	return ctx.Data
}

// Defense 角色进行防御，修改对应的伤害数据
func (c *Character) Defense(sender *Player, damage *AttackDamageContext) *DefenceDamageContext {
	damageData := NewDefenceDamageContext(sender, damage)
	ctx := NewContext(c.DefenceModifiers(), damageData)
	ctx.Continue()
	return ctx.Data
}

// Heal 对角色进行治疗
func (c *Character) Heal(amount uint) {
	if c.CurrentHealthPoint+amount < c.MaxHealthPoint {
		c.CurrentHealthPoint += amount
	} else {
		c.CurrentHealthPoint = c.MaxHealthPoint
	}

	c.callback(definition.AfterHeal)
}

// Charge 对角色进行充能
func (c *Character) Charge(amount uint) {
	if c.CurrentMagicPoint+amount < c.MaxMagicPoint {
		c.CurrentMagicPoint += amount
	} else {
		c.CurrentMagicPoint = c.MaxMagicPoint
	}

	c.callback(definition.AfterCharge)
}

// Background 将角色置于后台
func (c *Character) Background() {
	c.Status = definition.CharacterStatusBackground
}

// Active 将角色置于前台
func (c *Character) Active() {
	c.Status = definition.CharacterStatusActive

	c.callback(definition.AfterSwitch)
}

// Clone 将当前Character克隆一份
func (c *Character) Clone() *Character {
	return &Character{
		player: &Player{
			UID:              c.player.UID,
			DefenceModifiers: c.player.DefenceModifiers.Clone(),
		},
		ID:                 c.ID,
		Name:               c.Name,
		Title:              c.Title,
		Description:        c.Description,
		Affiliation:        c.Affiliation,
		VisionType:         c.VisionType,
		WeaponType:         c.WeaponType,
		Skills:             c.Skills,
		MaxHealthPoint:     c.MaxHealthPoint,
		CurrentHealthPoint: c.CurrentHealthPoint,
		MaxMagicPoint:      c.MaxMagicPoint,
		CurrentMagicPoint:  c.CurrentMagicPoint,
		ElementsAttached:   c.ElementsAttached,
		Equipments:         c.Equipments,
		Status:             c.Status,
		CostModifiers:      c.CostModifiers.Clone(),
		AttackModifiers:    c.AttackModifiers.Clone(),
		Satiety:            c.Satiety,
	}
}

func (c *Character) callback(trigger definition.Trigger) {
	data := NewCharacterContext(c)
	ctx := NewContext(c.CallbackChain[trigger], data)
	ctx.Continue()
}

// GenerateCharacter 通过外部实现的ICharacter创建运行时的Character实体
func GenerateCharacter(character ICharacter) (entity *Character) {
	// todo: complete all fields
	return &Character{
		ID:                 character.ID(),
		Name:               character.Name(),
		Title:              character.Title(),
		Description:        character.Description(),
		Affiliation:        character.Affiliation(),
		VisionType:         character.VisionType(),
		WeaponType:         character.WeaponType(),
		Skills:             character.Skills(),
		MaxHealthPoint:     character.MaxHealthPoint(),
		CurrentHealthPoint: character.MaxHealthPoint(),
		MaxMagicPoint:      character.MaxMagicPoint(),
		CurrentMagicPoint:  0,
		ElementsAttached:   []definition.Element{},
	}
}

type CharacterContext struct {
	character *Character
}

// ResetAttackModifiers 重置角色的攻击修正函数队列
func (ctx *CharacterContext) ResetAttackModifiers() {
	ctx.character.AttackModifiers = NewModifierChain[AttackDamageContext]()
}

// AddAttackModifier 为角色增加攻击修正函数
func (ctx *CharacterContext) AddAttackModifier(modifierName string, handler ModifierHandler[AttackDamageContext]) {
	ctx.character.AttackModifiers.AddModifierHandler(modifierName, handler)
}

// RemoveAttackModifier 为角色移除攻击修正函数
func (ctx *CharacterContext) RemoveAttackModifier(modifierName string) {
	ctx.character.AttackModifiers.RemoveModifierHandler(modifierName)
}

// ResetDefenceModifiers 重置角色的防御修正函数队列
func (ctx *CharacterContext) ResetDefenceModifiers() {
	ctx.character.player.DefenceModifiers = NewModifierChain[DefenceDamageContext]()
}

// AddDefenceModifier 为角色增加防御修正函数
func (ctx *CharacterContext) AddDefenceModifier(modifierName string, handler ModifierHandler[DefenceDamageContext]) {
	ctx.character.DefenceModifiers().AddModifierHandler(modifierName, handler)
}

// RemoveDefenceModifier 为角色移除防御修正函数
func (ctx *CharacterContext) RemoveDefenceModifier(modifierName string) {
	ctx.character.DefenceModifiers().RemoveModifierHandler(modifierName)
}

// SetSatiety 修改角色的饱腹状况
func (ctx *CharacterContext) SetSatiety(satiety bool) {
	ctx.character.Satiety = satiety
}

// Heal 对角色进行治疗
func (ctx *CharacterContext) Heal(amount uint) {
	ctx.character.Heal(amount)
}

// Charge 对角色进行充能
func (ctx *CharacterContext) Charge(amount uint) {
	ctx.character.Charge(amount)
}

// SwitchPrevCharacter 将前台角色切换到前一个角色，如果可能
func (ctx *CharacterContext) SwitchPrevCharacter() {
	canSwitch := false
	for _, character := range ctx.character.player.Characters {
		if character.Status != definition.CharacterStatusUnselectable && character != ctx.character.player.ActiveCharacter {
			canSwitch = true
			break
		}
	}

	if canSwitch {
		index := 0
		for i := 0; i < len(ctx.character.player.Characters); i++ {
			if ctx.character.player.Characters[i] == ctx.character.player.ActiveCharacter {
				index = i
				break
			}
		}

		for i := index - 1; i >= 0; i-- {
			if character := ctx.character.player.Characters[i]; character.Status != definition.CharacterStatusUnselectable {
				ctx.SwitchCharacter(character)
				return
			}
		}

		for i := len(ctx.character.player.Characters) - 1; i > index; i-- {
			if character := ctx.character.player.Characters[i]; character.Status != definition.CharacterStatusUnselectable {
				ctx.SwitchCharacter(character)
				return
			}
		}
	}
}

// SwitchNextCharacter 将前台角色切换到后一个角色，如果可能
func (ctx *CharacterContext) SwitchNextCharacter() {
	canSwitch := false
	for _, character := range ctx.character.player.Characters {
		if character.Status != definition.CharacterStatusUnselectable && character != ctx.character.player.ActiveCharacter {
			canSwitch = true
			break
		}
	}

	if canSwitch {
		index := 0
		for i := 0; i < len(ctx.character.player.Characters); i++ {
			if ctx.character.player.Characters[i] == ctx.character.player.ActiveCharacter {
				index = i
				break
			}
		}

		for i := index + 1; i < len(ctx.character.player.Characters); i++ {
			if character := ctx.character.player.Characters[i]; character.Status != definition.CharacterStatusUnselectable {
				ctx.SwitchCharacter(character)
				return
			}
		}

		for i := 0; i < index; i++ {
			if character := ctx.character.player.Characters[i]; character.Status != definition.CharacterStatusUnselectable {
				ctx.SwitchCharacter(character)
				return
			}
		}
	}
}

// SwitchCharacter 将前台角色切换为指定角色，如果可能
func (ctx *CharacterContext) SwitchCharacter(character *Character) {
	if ctx.character.player.ActiveCharacter != character {
		ctx.character.player.ActiveCharacter.Background()
		ctx.character.player.ActiveCharacter = character
		character.Active()
	}
}

// NewCharacterContext 使用Character生成一个CharacterContext
func NewCharacterContext(character *Character) *CharacterContext {
	return &CharacterContext{character: character}
}
