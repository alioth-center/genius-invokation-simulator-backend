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
	activeDamage     func(ctx definition.Context) definition.Damage
	backgroundDamage func(ctx definition.Context) definition.Damage
}

func (impl *AttackSkillImpl) SkillCost() map[enum.ElementType]uint {
	return impl.skillCost
}

func (impl *AttackSkillImpl) ActiveDamage(ctx definition.Context) definition.Damage {
	return impl.activeDamage(ctx)
}

func (impl *AttackSkillImpl) BackgroundDamage(ctx definition.Context) definition.Damage {
	return impl.backgroundDamage(ctx)
}

type AttackSkillOptions func(option *AttackSkillImpl)

func WithAttackSkillID(id uint16) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		opt := WithSkillID(id)
		opt(&option.SkillImpl)
	}
}

func WithAttackSkillType(skillType enum.SkillType) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		opt := WithSkillType(skillType)
		opt(&option.SkillImpl)
	}
}

func WithAttackSkillCost(skillCost map[enum.ElementType]uint) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.skillCost = skillCost
	}
}

func WithAttackSkillActiveDamageHandler(handler func(ctx definition.Context) (elementType enum.ElementType, damageAmount uint)) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.activeDamage = func(ctx definition.Context) definition.Damage {
			elementType, damageAmount := handler(ctx)
			return NewDamageWithOpts(
				WithDamageElementType(elementType),
				WithDamageAmount(damageAmount),
			)
		}
	}
}

func WithAttackSkillBackgroundDamageHandler(handler func(ctx definition.Context) (damageAmount uint)) AttackSkillOptions {
	return func(option *AttackSkillImpl) {
		option.backgroundDamage = func(ctx definition.Context) definition.Damage {
			damageAmount := handler(ctx)
			return NewDamageWithOpts(
				WithDamageElementType(enum.ElementNone),
				WithDamageAmount(damageAmount),
			)
		}
	}
}

func NewAttackSkillWithOpts(options ...AttackSkillOptions) definition.AttackSkill {
	impl := &AttackSkillImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}
