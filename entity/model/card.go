package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Card interface {
	BaseEntity
	Type() enum.CardType
	Cost() map[enum.ElementType]uint
}

type EventCard interface {
	Card
	Event() (event Event)
}

type SupportCard interface {
	Card
	Support() (event Event)
}

type EquipmentCard interface {
	Card
	EquipmentType() enum.EquipmentType
	Modify() (event Event)
}

type WeaponCard interface {
	EquipmentCard
	WeaponType() enum.WeaponType
}

func ConvertToEventCard(original Card) (success bool, result EventCard) {
	cardType := original.Type()
	if cardType != enum.CardEvent && cardType != enum.CardElementalResonance && cardType != enum.CardFood {
		return false, result
	} else if convertedCard, convertSuccess := original.(EventCard); !convertSuccess {
		return false, result
	} else {
		return true, convertedCard
	}
}

func ConvertToSupportCard(original Card) (success bool, result SupportCard) {
	cardType := original.Type()
	if cardType != enum.CardSupport && cardType != enum.CardCompanion && cardType != enum.CardLocation && cardType != enum.CardItem {
		return false, result
	} else if convertedCard, convertSuccess := original.(SupportCard); !convertSuccess {
		return false, result
	} else {
		return true, convertedCard
	}
}

func ConvertToEquipmentCard(original Card) (success bool, result EquipmentCard) {
	cardType := original.Type()
	if cardType != enum.CardEquipment && cardType != enum.CardTalent && cardType != enum.CardWeapon && cardType != enum.CardArtifact {
		return false, result
	} else if convertedCard, convertSuccess := original.(EquipmentCard); !convertSuccess {
		return false, result
	} else {
		return true, convertedCard
	}
}
