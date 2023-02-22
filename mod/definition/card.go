package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Card interface {
	model.BaseEntity
	CardType() enum.CardType
	Cost() map[enum.ElementType]uint
}

type EventCard interface {
	Card
	Event() Event
}

type SupportCard interface {
	Card
	Support() Event
}

type EquipmentCard interface {
	Card
	EquipmentType() enum.EquipmentType
	Modify() Event
}

type WeaponCard interface {
	EquipmentCard
	WeaponType() enum.WeaponType
}
