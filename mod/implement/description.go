package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type DescriptionImpl struct {
	descriptionType enum.DescriptionType
	shortID         uint16
	longID          uint64
}

func (impl *DescriptionImpl) DescriptionType() enum.DescriptionType {
	return impl.descriptionType
}

func (impl *DescriptionImpl) ShortID() uint16 {
	return impl.shortID
}

func (impl *DescriptionImpl) LongID() uint64 {
	return impl.longID
}

type CharacterDescriptionImpl struct {
	DescriptionImpl
	characterName        string
	characterDescription string
	characterTitle       string
	characterStory       string
}

func (impl *CharacterDescriptionImpl) CharacterName() string {
	return impl.characterName
}

func (impl *CharacterDescriptionImpl) CharacterDescription() string {
	return impl.characterDescription
}

func (impl *CharacterDescriptionImpl) CharacterTitle() string {
	return impl.characterTitle
}

func (impl *CharacterDescriptionImpl) CharacterStory() string {
	return impl.characterStory
}

type SkillDescriptionImpl struct {
	DescriptionImpl
	skillName        string
	skillDescription string
}

func (impl *SkillDescriptionImpl) SkillName() string {
	return impl.skillName
}

func (impl *SkillDescriptionImpl) SkillDescription() string {
	return impl.skillDescription
}

type EventDescriptionImpl struct {
	DescriptionImpl
	eventName        string
	eventDescription string
}

func (impl *EventDescriptionImpl) EventName() string {
	return impl.eventName
}

func (impl *EventDescriptionImpl) EventDescription() string {
	return impl.eventDescription
}

type CardDescriptionImpl struct {
	DescriptionImpl
	cardName        string
	cardDescription string
}

func (impl *CardDescriptionImpl) CardName() string {
	return impl.cardName
}

func (impl *CardDescriptionImpl) CardDescription() string {
	return impl.cardDescription
}

type ModifierDescriptionImpl struct {
	DescriptionImpl
	modifierName        string
	modifierDescription string
}

func (impl *ModifierDescriptionImpl) ModifierName() string {
	return impl.modifierName
}

func (impl *ModifierDescriptionImpl) ModifierDescription() string {
	return impl.modifierDescription
}

type ModDescriptionImpl struct {
	id          uint64
	name        string
	description string
}

func (impl *ModDescriptionImpl) ModID() uint64 {
	return impl.id
}

func (impl *ModDescriptionImpl) ModName() string {
	return impl.name
}

func (impl *ModDescriptionImpl) ModDescription() string {
	return impl.description
}
