package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
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

type DescriptionOptions func(option *DescriptionImpl)

func WithDescriptionType(descriptionType enum.DescriptionType) DescriptionOptions {
	return func(option *DescriptionImpl) {
		option.descriptionType = descriptionType
	}
}

func WithDescriptionID(id uint16) DescriptionOptions {
	return func(option *DescriptionImpl) {
		entity := NewEntityWithOpts(WithEntityID(id))
		option.shortID = id
		option.longID = entity.TypeID()
	}
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

type CharacterDescriptionOptions func(option *CharacterDescriptionImpl)

func WithCharacterDescriptionType(descriptionType enum.DescriptionType) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		opt := WithDescriptionType(descriptionType)
		opt(&option.DescriptionImpl)
	}
}

func WithCharacterDescriptionID(id uint16) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithCharacterDescriptionName(name string) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		option.characterName = name
	}
}

func WithCharacterDescriptionDescription(description string) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		option.characterDescription = description
	}
}

func WithCharacterDescriptionTitle(title string) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		option.characterTitle = title
	}
}

func WithCharacterDescriptionStory(story string) CharacterDescriptionOptions {
	return func(option *CharacterDescriptionImpl) {
		option.characterStory = story
	}
}

func NewCharacterDescriptionWithOpts(options ...CharacterDescriptionOptions) definition.CharacterDescription {
	impl := &CharacterDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
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

type SkillDescriptionOptions func(option *SkillDescriptionImpl)

func WithSkillDescriptionType(descriptionType enum.DescriptionType) SkillDescriptionOptions {
	return func(option *SkillDescriptionImpl) {
		opt := WithDescriptionType(descriptionType)
		opt(&option.DescriptionImpl)
	}
}

func WithSkillDescriptionID(id uint16) SkillDescriptionOptions {
	return func(option *SkillDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithSkillDescriptionName(name string) SkillDescriptionOptions {
	return func(option *SkillDescriptionImpl) {
		option.skillName = name
	}
}

func WithSkillDescriptionDescription(description string) SkillDescriptionOptions {
	return func(option *SkillDescriptionImpl) {
		option.skillDescription = description
	}
}

func NewSkillDescriptionWithOpts(options ...SkillDescriptionOptions) definition.SkillDescription {
	impl := &SkillDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
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

type EventDescriptionOptions func(option *EventDescriptionImpl)

func WithEventDescriptionType(descriptionType enum.DescriptionType) EventDescriptionOptions {
	return func(option *EventDescriptionImpl) {
		opt := WithDescriptionType(descriptionType)
		opt(&option.DescriptionImpl)
	}
}

func WithEventDescriptionID(id uint16) EventDescriptionOptions {
	return func(option *EventDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithEventDescriptionName(name string) EventDescriptionOptions {
	return func(option *EventDescriptionImpl) {
		option.eventName = name
	}
}

func WithEventDescriptionDescription(description string) EventDescriptionOptions {
	return func(option *EventDescriptionImpl) {
		option.eventDescription = description
	}
}

func NewEventDescriptionWithOpts(options ...EventDescriptionOptions) definition.EventDescription {
	impl := &EventDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
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

type CardDescriptionOptions func(option *CardDescriptionImpl)

func WithCardDescriptionType(descriptionType enum.DescriptionType) CardDescriptionOptions {
	return func(option *CardDescriptionImpl) {
		opt := WithDescriptionType(descriptionType)
		opt(&option.DescriptionImpl)
	}
}

func WithCardDescriptionID(id uint16) CardDescriptionOptions {
	return func(option *CardDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithCardDescriptionName(name string) CardDescriptionOptions {
	return func(option *CardDescriptionImpl) {
		option.cardName = name
	}
}

func WithCardDescriptionDescription(description string) CardDescriptionOptions {
	return func(option *CardDescriptionImpl) {
		option.cardDescription = description
	}
}

func NewCardDescriptionWithOpts(options ...CardDescriptionOptions) definition.CardDescription {
	impl := &CardDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}

type SummonDescriptionImpl struct {
	DescriptionImpl
	summonName        string
	summonDescription string
}

func (impl *SummonDescriptionImpl) SummonName() string {
	return impl.summonName
}

func (impl *SummonDescriptionImpl) SummonDescription() string {
	return impl.summonDescription
}

type SummonDescriptionOptions func(option *SummonDescriptionImpl)

func WithSummonDescriptionID(id uint16) SummonDescriptionOptions {
	return func(option *SummonDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithSummonDescriptionName(name string) SummonDescriptionOptions {
	return func(option *SummonDescriptionImpl) {
		option.summonName = name
	}
}

func WithSummonDescriptionDescription(description string) SummonDescriptionOptions {
	return func(option *SummonDescriptionImpl) {
		option.summonDescription = description
	}
}

func NewSummonDescriptionWithOpts(options ...SummonDescriptionOptions) definition.SummonDescription {
	impl := &SummonDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
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

type ModifierDescriptionOptions func(option *ModifierDescriptionImpl)

func WithModifierDescriptionType(descriptionType enum.DescriptionType) ModifierDescriptionOptions {
	return func(option *ModifierDescriptionImpl) {
		opt := WithDescriptionType(descriptionType)
		opt(&option.DescriptionImpl)
	}
}

func WithModifierDescriptionID(id uint16) ModifierDescriptionOptions {
	return func(option *ModifierDescriptionImpl) {
		opt := WithDescriptionID(id)
		opt(&option.DescriptionImpl)
	}
}

func WithModifierDescriptionName(name string) ModifierDescriptionOptions {
	return func(option *ModifierDescriptionImpl) {
		option.modifierName = name
	}
}

func WithModifierDescriptionDescription(description string) ModifierDescriptionOptions {
	return func(option *ModifierDescriptionImpl) {
		option.modifierDescription = description
	}
}

func NewModifierDescriptionWithOpts(options ...ModifierDescriptionOptions) definition.ModifierDescription {
	impl := &ModifierDescriptionImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
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

type ModDescriptionOptions func(option *ModDescriptionImpl)

func WithModName(name string) ModDescriptionOptions {
	return func(option *ModDescriptionImpl) {
		option.name = name
	}
}

func WithModDescription(description string) ModDescriptionOptions {
	return func(option *ModDescriptionImpl) {
		option.description = description
	}
}

func NewModDescriptionWithOpts(options ...ModDescriptionOptions) definition.ModDescription {
	impl := &ModDescriptionImpl{
		id: ModID(),
	}

	for _, option := range options {
		option(impl)
	}

	return impl
}
