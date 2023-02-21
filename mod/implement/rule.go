package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

var necessaryImplements = []enum.RuleType{
	enum.RuleTypeReactionCalculator,
	enum.RuleTypeVictorCalculator,
}

func RuleConvert[dest any](source interface{}) (success bool, result dest) {
	converted, ok := source.(dest)
	return ok, converted
}

type RuleImpl struct {
	EntityImpl
	ruleImpl map[enum.RuleType]interface{}
}

func (r *RuleImpl) CopyFrom(source definition.Rule, filter ...enum.RuleType) {
	if r.ruleImpl == nil {
		r.ruleImpl = map[enum.RuleType]interface{}{}
	}

	for _, ruleType := range filter {
		r.ruleImpl[ruleType] = source.Implements(ruleType)
	}
}

func (r *RuleImpl) Implements(ruleType enum.RuleType) interface{} {
	if r.ruleImpl == nil {
		r.ruleImpl = map[enum.RuleType]interface{}{}
	}

	return r.ruleImpl[ruleType]
}

func (r *RuleImpl) CheckImplements() (success bool) {
	if r.ruleImpl == nil {
		return false
	}

	if reactionCalculator, has := r.ruleImpl[enum.RuleTypeReactionCalculator]; !has || reactionCalculator == nil {
		return false
	} else if success, _ := RuleConvert[model.ReactionCalculator](reactionCalculator); !success {
		return false
	}

	if victorCalculator, has := r.ruleImpl[enum.RuleTypeVictorCalculator]; !has || victorCalculator == nil {
		return false
	} else if success, _ := RuleConvert[model.VictorCalculator](victorCalculator); !success {
		return false
	}

	return true
}

type RuleImplOptions func(option *RuleImpl)

func WithRuleID(id uint16) RuleImplOptions {
	return func(option *RuleImpl) {
		option.EntityImpl.InjectTypeID(uint64(id))
	}
}

func WithRuleImplement(ruleType enum.RuleType, implement interface{}) RuleImplOptions {
	return func(option *RuleImpl) {
		if option.ruleImpl == nil {
			option.ruleImpl = map[enum.RuleType]interface{}{}
		}

		option.ruleImpl[ruleType] = implement
	}
}

func WithRuleCopyFrom(another definition.Rule, filter ...enum.RuleType) RuleImplOptions {
	return func(option *RuleImpl) {
		if option.ruleImpl == nil {
			option.ruleImpl = map[enum.RuleType]interface{}{}
		}

		for _, ruleType := range filter {
			option.ruleImpl[ruleType] = another.Implements(ruleType)
		}
	}
}

func NewRuleWithOpts(options ...RuleImplOptions) definition.Rule {
	impl := &RuleImpl{
		ruleImpl: map[enum.RuleType]interface{}{},
	}

	for _, option := range options {
		option(impl)
	}

	if !impl.CheckImplements() {
		// 未完全实现规则集合，避免在运行时抛出panic，在初始化时抛出
		panic("implement check failed")
	}

	return impl
}
