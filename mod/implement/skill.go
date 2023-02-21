package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type SkillImpl struct {
	EntityImpl
	skillType enum.SkillType
}

func (impl *SkillImpl) SkillType() enum.SkillType {
	return impl.skillType
}

type SkillOptions func(option *SkillImpl)

func WithSkillID(id uint16) SkillOptions {
	return func(option *SkillImpl) {
		option.InjectTypeID(uint64(id))
	}
}

func WithSkillType(skillType enum.SkillType) SkillOptions {
	return func(option *SkillImpl) {
		option.skillType = skillType
	}
}

func NewSkillWithOpts(options ...SkillOptions) definition.Skill {
	impl := &SkillImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}

type AttackSkillImpl struct {
	SkillImpl
	skillCost        map[enum.ElementType]uint
	activeDamage     definition.Damage
	backgroundDamage definition.Damage
}

func (impl *AttackSkillImpl) SkillCost() map[enum.ElementType]uint {
	return impl.skillCost
}

func (impl *AttackSkillImpl) ActiveDamage() definition.Damage {
	return impl.activeDamage
}

func (impl *AttackSkillImpl) BackgroundDamage() definition.Damage {
	return impl.backgroundDamage
}

type AttackSkillOptions func(option *AttackSkillImpl)

func WithAttackSkillID(id uint16) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		opt := WithSkillID(id)
		opt(&option.SkillImpl)
	}
}

func WithAttackSkillCost(skillCost map[enum.ElementType]uint) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.skillCost = skillCost
	}
}

func WithAttackSkillActiveDamage(elementType enum.ElementType, damageAmount uint) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.activeDamage = NewDamageWithOpts(
			WithDamageElementType(elementType),
			WithDamageAmount(damageAmount),
		)
	}
}

func WithAttackSkillBackgroundDamage(damageAmount uint) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.backgroundDamage = NewDamageWithOpts(
			WithDamageElementType(enum.ElementNone),
			WithDamageAmount(damageAmount),
		)
	}
}

func NewAttackSkillWithOpts(options ...AttackSkillOptions) definition.AttackSkill {
	impl := &AttackSkillImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}
