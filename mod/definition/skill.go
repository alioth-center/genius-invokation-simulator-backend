package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Skill interface {
	model.BaseEntity
	SkillType() enum.SkillType
}

type AttackSkill interface {
	Skill
	SkillCost() map[enum.ElementType]uint
	ActiveDamage(ctx Context) Damage
	BackgroundDamage(ctx Context) Damage
}

type CooperativeSkill interface {
	Skill
	ActiveDamage(ctx Context) Damage
	BackgroundDamage(ctx Context) Damage
	EffectLeft(ctx Context) uint
	Effective(ctx Context) bool
}
