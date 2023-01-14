package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type BaseCardImplement struct {
	BaseEntityImplement
}

type EventCardImplement struct {
	BaseCardImplement
}

func (implement *EventCardImplement) Type() enum.CardType {
	return enum.CardEvent
}

type FoodCardImplement struct {
	EventCardImplement
}

func (implement *FoodCardImplement) Type() enum.CardType {
	return enum.CardFood
}

type ElementalResonanceCardImplement struct {
	EventCardImplement
}

func (implement *ElementalResonanceCardImplement) Type() enum.CardType {
	return enum.CardElementalResonance
}

type EquipmentCardImplement struct {
	BaseCardImplement
}

func (implement *EquipmentCardImplement) Cost() map[enum.ElementType]uint {
	//TODO implement me
	panic("implement me")
}

func (implement *EquipmentCardImplement) EquipmentType() enum.EquipmentType {
	//TODO implement me
	panic("implement me")
}

func (implement *EquipmentCardImplement) Modify() (event model.Event) {
	//TODO implement me
	panic("implement me")
}

func (implement *EquipmentCardImplement) Type() enum.CardType {
	return enum.CardEquipment
}

type TalentCardImplement struct {
	EquipmentCardImplement
}

func (implement *TalentCardImplement) Type() enum.CardType {
	return enum.CardTalent
}

type WeaponCardImplement struct {
	EquipmentCardImplement
}

func (implement *WeaponCardImplement) Type() enum.CardType {
	return enum.CardWeapon
}

type ArtifactCardImplement struct {
	EquipmentCardImplement
}

func (implement *ArtifactCardImplement) Type() enum.CardType {
	return enum.CardArtifact
}
