package adapter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
)

type RuleSetAdapter struct{}

func (r RuleSetAdapter) Convert(source definition.Rule) (success bool, result model.RuleSet) {
	adapterLayer := model.RuleSet{}
	adapterLayer.ID = source.TypeID()
	_, adapterLayer.ReactionCalculator = implement.RuleConvert[model.ReactionCalculator](source.Implements(enum.RuleTypeReactionCalculator))
	_, adapterLayer.VictorCalculator = implement.RuleConvert[model.VictorCalculator](source.Implements(enum.RuleTypeVictorCalculator))

	return true, adapterLayer
}

func NewRuleSetAdapter() adapter.Adapter[definition.Rule, model.RuleSet] {
	return RuleSetAdapter{}
}
