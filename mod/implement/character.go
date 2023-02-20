package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type CharacterImpl struct {
	name        string
	affiliation enum.Affiliation
	vision      enum.ElementType
	weapon      enum.WeaponType
	skills      []definition.Skill
	hp          uint
	mp          uint
}

func (c CharacterImpl) Name() string {
	return c.name
}

func (c CharacterImpl) Affiliation() enum.Affiliation {
	return c.affiliation
}

func (c CharacterImpl) Vision() enum.ElementType {
	return c.vision
}

func (c CharacterImpl) Weapon() enum.WeaponType {
	return c.weapon
}

func (c CharacterImpl) Skills() []definition.Skill {
	return c.skills
}

func (c CharacterImpl) HP() uint {
	return c.hp
}

func (c CharacterImpl) MP() uint {
	return c.mp
}

func WithCharacterName(name string) CharacterOptions {
	return func(option *CharacterImpl) {
		option.name = name
	}
}

func WithCharacterAffiliation(affiliation enum.Affiliation) CharacterOptions {
	return func(option *CharacterImpl) {
		option.affiliation = affiliation
	}
}

func WithCharacterVision(vision enum.ElementType) CharacterOptions {
	return func(option *CharacterImpl) {
		option.vision = vision
	}
}

func WithCharacterWeapon(weapon enum.WeaponType) CharacterOptions {
	return func(option *CharacterImpl) {
		option.weapon = weapon
	}
}

func WithCharacterSkills(skills ...definition.Skill) CharacterOptions {
	return func(option *CharacterImpl) {
		option.skills = skills
	}
}

func WithCharacterHP(hp uint) CharacterOptions {
	return func(option *CharacterImpl) {
		option.hp = hp
	}
}

func WithCharacterMP(mp uint) CharacterOptions {
	return func(option *CharacterImpl) {
		option.mp = mp
	}
}

type CharacterOptions func(option *CharacterImpl)

func NewCharacterWithOpts(opts ...CharacterOptions) definition.Character {
	character := &CharacterImpl{}
	for _, opt := range opts {
		opt(character)
	}

	return character
}
