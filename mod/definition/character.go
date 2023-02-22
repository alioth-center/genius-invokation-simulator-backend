package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Character interface {
	model.BaseEntity
	Name() string
	Affiliation() enum.Affiliation
	Vision() enum.ElementType
	Weapon() enum.WeaponType
	Skills() []Skill
	HP() uint
	MP() uint
}
