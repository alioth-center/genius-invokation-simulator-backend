package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Skill struct {
	PersistenceIndex[entity.Skill]
	Type enum.SkillType
}
