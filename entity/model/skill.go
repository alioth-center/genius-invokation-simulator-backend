package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Skill interface {
	ID() uint
	Type() enum.SkillType
}

type AttackSkill interface {
	Skill
	Cost() Cost
	BaseDamage(targetCharacter, senderCharacter uint, backgroundCharacters []uint) *context.DamageContext
}

type CooperativeSkill interface {
	Skill
	BaseDamage() *context.DamageContext
	EffectLeft() uint
	Effective() bool
}

type PassiveSkill interface {
}
