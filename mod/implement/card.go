package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	definition "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	definition2 "github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type BaseCardImpl struct {
}

func (b *BaseCardImpl) CardType() enum.CardType {
	return enum.CardUnknown
}

func (b *BaseCardImpl) Cost() map[enum.ElementType]uint {
	return nil
}

type EventCardImpl struct {
	BaseCardImpl
}

func (e *EventCardImpl) CardType() enum.CardType {
	return enum.CardEvent
}

func (e *EventCardImpl) Event() definition.Event {
	return nil
}

type FoodCardImpl struct {
	EventCardImpl
}

func (e *FoodCardImpl) CardType() enum.CardType {
	return enum.CardFood
}

type ElementalResonanceCardImpl struct {
	EventCardImpl
}

func (e *ElementalResonanceCardImpl) CardType() enum.CardType {
	return enum.CardElementalResonance
}

type EquipmentCardImpl struct {
	BaseCardImpl
}

func (e *EquipmentCardImpl) CardType() enum.CardType {
	return enum.CardEquipment
}

func (e *EquipmentCardImpl) EquipmentType() enum.EquipmentType {
	return enum.EquipmentNone
}

func (e *EquipmentCardImpl) Modify() definition.Event {
	return nil
}

type ArtifactCardImpl struct {
	EquipmentCardImpl
}

func (e *ArtifactCardImpl) CardType() enum.CardType {
	return enum.CardArtifact
}

func (e *ArtifactCardImpl) EquipmentType() enum.EquipmentType {
	return enum.EquipmentArtifact
}

type TalentCardImpl struct {
	EquipmentCardImpl
}

func (e *TalentCardImpl) CardType() enum.CardType {
	return enum.CardTalent
}

func (e *TalentCardImpl) EquipmentType() enum.EquipmentType {
	return enum.EquipmentTalent
}

type WeaponCardImpl struct {
	EquipmentCardImpl
}

func (e *WeaponCardImpl) WeaponType() enum.WeaponType {
	return enum.WeaponOthers
}

func (e *WeaponCardImpl) CardType() enum.CardType {
	return enum.CardWeapon
}

func (e *WeaponCardImpl) EquipmentType() enum.EquipmentType {
	return enum.EquipmentWeapon
}

type SkyBow struct {
	EquipmentCardImpl
}

func (s *SkyBow) Cost() map[enum.ElementType]uint {
	return map[enum.ElementType]uint{
		enum.ElementSame: 3,
	}
}

func (s *SkyBow) WeaponType() enum.WeaponType {
	return enum.WeaponBow
}

func SkyBowModifier() definition2.Modifier[context.DamageContext] {
	panic("")
}
