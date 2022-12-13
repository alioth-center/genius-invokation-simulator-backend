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
	elementFilter      func(elements []definition.Element) (result []definition.Element)
}

func (character *Character) OnAttachedElement(element definition.Element) {

}

func (character *Character) OnAttacked(attackElement definition.Element, damage uint) {

}

func GenerateCharacter(character ICharacter) (entity *Character) {
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
		elementFilter:      character.ElementFilter,
	}
}
