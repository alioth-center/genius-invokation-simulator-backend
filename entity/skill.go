package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Skill interface {
	ID() uint
	Name() string
	Type() enum.SkillType
}

type AttackSkill interface {
	Cost()
	BaseDamage() *context.DamageContext
}

type CooperativeSkill interface {
	BaseDamage() *context.DamageContext
	EffectLeft() uint
	Effective() bool
}

type PassiveSkill interface {
}
