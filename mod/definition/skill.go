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
	ActiveDamage() Damage
	BackgroundDamage() Damage
}

type CooperativeSkill interface {
	Skill
	ActiveDamage() Damage
	BackgroundDamage() Damage
	EffectLeft() uint
	Effective() bool
}
