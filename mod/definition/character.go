package definition

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type Character interface {
	Name() string
	Affiliation() enum.Affiliation
	Vision() enum.ElementType
	Weapon() enum.WeaponType
	Skills() []Skill
	HP() uint
	MP() uint
}

type characterInfo struct {
	name        string
	affiliation enum.Affiliation
	vision      enum.ElementType
	weapon      enum.WeaponType
	skills      []Skill
	hp          uint
	mp          uint
}

func (c characterInfo) Name() string {
	return c.name
}

func (c characterInfo) Affiliation() enum.Affiliation {
	return c.affiliation
}

func (c characterInfo) Vision() enum.ElementType {
	return c.vision
}

func (c characterInfo) Weapon() enum.WeaponType {
	return c.weapon
}

func (c characterInfo) Skills() []Skill {
	return c.skills
}

func (c characterInfo) HP() uint {
	return c.hp
}

func (c characterInfo) MP() uint {
	return c.mp
}

func WithCharacterName(name string) CharacterOptions {
	return func(option *characterInfo) {
		option.name = name
	}
}

func WithCharacterAffiliation(affiliation enum.Affiliation) CharacterOptions {
	return func(option *characterInfo) {
		option.affiliation = affiliation
	}
}

func WithCharacterVision(vision enum.ElementType) CharacterOptions {
	return func(option *characterInfo) {
		option.vision = vision
	}
}

func WithCharacterWeapon(weapon enum.WeaponType) CharacterOptions {
	return func(option *characterInfo) {
		option.weapon = weapon
	}
}

func WithCharacterSkills(skills ...Skill) CharacterOptions {
	return func(option *characterInfo) {
		option.skills = skills
	}
}

func WithCharacterHP(hp uint) CharacterOptions {
	return func(option *characterInfo) {
		option.hp = hp
	}
}

func WithCharacterMP(mp uint) CharacterOptions {
	return func(option *characterInfo) {
		option.mp = mp
	}
}

type CharacterOptions func(option *characterInfo)

func NewCharacterWithOpts(opts ...CharacterOptions) Character {
	character := &characterInfo{}
	for _, opt := range opts {
		opt(character)
	}

	return character
}
