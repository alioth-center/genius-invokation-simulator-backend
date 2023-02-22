package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Rule interface {
	model.BaseEntity
	CopyFrom(source Rule, filter ...enum.RuleType)
	Implements(ruleType enum.RuleType) interface{}
	CheckImplements() (success bool)
}
