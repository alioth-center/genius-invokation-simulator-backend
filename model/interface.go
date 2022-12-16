/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "interface.go" LastUpdatedAt 2022/12/12 15:54:12
 */

package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
)

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

type IConsumable interface {
	Cost() definition.ElementSet
}

type ICard interface {
	ID() uint
	Name() string
	Description() string
	Type() definition.CardType
	Cost() definition.ElementSet
}

type IEquipmentCard interface {
	ICard
	EquipmentType() definition.EquipmentType
	Equip() ModifierHandler[CharacterContext]
}

type IEventCard interface {
	Event() ModifierHandler[CharacterContext]
	CallBack() (callbackTrigger definition.Trigger, callbackEvent ICallbackEvent)
}

type ICallbackEvent interface {
	Name() string
	Event() ModifierHandler[CharacterContext]
	Triggered() bool
}

type ISkill interface {
	Name() string
	Description() string
	Type() definition.SkillType
	Buffer() ModifierHandler[CharacterContext]
}

type IAttackSkill interface {
	ISkill
	Attack(target *Player) *AttackDamageContext
}

type INormalSkill interface {
	IAttackSkill
	Cost() definition.ElementSet
}

type IPassiveSkill interface {
	ISkill
	Cost() definition.ElementSet
	HandlerFunc() ModifierHandler[CharacterContext]
}

type ICooperativeSkill interface {
	IAttackSkill
	Effective() bool
}
