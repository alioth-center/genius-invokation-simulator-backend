/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "interface.go" LastUpdatedAt 2022/12/12 15:54:12
 */

package model

import "github.com/sunist-c/genius-invokation-simulator-backend/definition"

type ICharacter interface {
	ID() uint
	Name() string
	Title() string
	Description() string
	Affiliation() definition.Affiliation
	VisionType() definition.Element
	WeaponType() definition.Weapon
	Skills() []ISkill
	MaxHealthPoint() uint
	MaxMagicPoint() uint
}

type ICard interface {
	ID() uint
	Name() string
	Description() string
	Type() definition.CardType
	Cost() definition.ElementSet
}

type ISkill interface {
	Name() string
	Description() string
	Cost() definition.ElementSet
	Type() definition.SkillType
	Buffer() func(self *Player)
}

type INormalAttack interface {
	ISkill
	NormalAttack(target *Player) AttackDamageContext
}

type IElementalSkill interface {
	ISkill
	ElementalAttack(target *Player) AttackDamageContext
}

type IElementalBurst interface {
	ISkill
	BurstAttack(target *Player) AttackDamageContext
}

type IPassiveSkill interface {
	ISkill
	HandlerFunc() ModifierHandler[any]
}
