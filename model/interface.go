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
	CurrentHealthPoint() uint
	MaxMagicPoint() uint
	CurrentMagicPoint() uint
	Status() definition.CharacterStatus

	Initialize()
}

type ICard interface {
	ID() uint
	Name() string
	Description() string
	Type() definition.CardType
	Cost() definition.ElementSet
}

type ISkill interface {
	ID() uint
	Name() string
	Description() string
	Cost() definition.ElementSet
	Type() definition.SkillType
}
