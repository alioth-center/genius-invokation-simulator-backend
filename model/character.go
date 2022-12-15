/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "character.go" LastUpdatedAt 2022/12/12 17:41:12
 */

package model

import "github.com/sunist-c/genius-invokation-simulator-backend/definition"

type Character struct {
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
	DefenceModifiers   ModifierChain[DefenceDamageContext]
	CostModifiers      ModifierChain[CostContext]
	AttackModifiers    ModifierChain[AttackDamageContext]
	HealModifiers      ModifierChain[HealContext]
	Satiety            bool
}

// Attack 角色释放攻击，生成相应的伤害数据
func (character *Character) Attack(target *Player, skill IAttackSkill) *AttackDamageContext {
	damageData := skill.Attack(target)
	ctx := NewContext(character.AttackModifiers, damageData)
	ctx.Continue()
	return ctx.Data
}

// Defense 角色进行防御，修改对应的伤害数据
func (character *Character) Defense(sender *Player, damage *AttackDamageContext) *DefenceDamageContext {
	damageData := NewDefenceDamageContext(sender, damage)
	ctx := NewContext(character.DefenceModifiers, damageData)
	ctx.Continue()
	return ctx.Data
}

// Heal 对角色进行治疗
func (character *Character) Heal(amount uint) *HealContext {
	healData := NewHealContext(amount)
	ctx := NewContext(character.HealModifiers, healData)
	ctx.Continue()
	return ctx.Data
}

// Clone 将当前Character克隆一份
func (character *Character) Clone() *Character {
	return &Character{
		ID:                 character.ID,
		Name:               character.Name,
		Title:              character.Title,
		Description:        character.Description,
		Affiliation:        character.Affiliation,
		VisionType:         character.VisionType,
		WeaponType:         character.WeaponType,
		Skills:             character.Skills,
		MaxHealthPoint:     character.MaxHealthPoint,
		CurrentHealthPoint: character.CurrentHealthPoint,
		MaxMagicPoint:      character.MaxMagicPoint,
		CurrentMagicPoint:  character.CurrentMagicPoint,
		ElementsAttached:   character.ElementsAttached,
		Equipments:         character.Equipments,
		Status:             character.Status,
		DefenceModifiers:   character.DefenceModifiers,
		CostModifiers:      character.CostModifiers,
		AttackModifiers:    character.AttackModifiers,
		HealModifiers:      character.HealModifiers,
		Satiety:            character.Satiety,
	}
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
