package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type LanguagePackImpl struct {
	language             enum.Language
	modDescription       definition.ModDescription
	characterDescription map[uint64]definition.CharacterDescription
	skillDescription     map[uint64]definition.SkillDescription
	eventDescription     map[uint64]definition.EventDescription
	cardDescription      map[uint64]definition.CardDescription
	modifierDescription  map[uint64]definition.ModifierDescription
}

func (impl *LanguagePackImpl) Language() enum.Language {
	return impl.language
}

func (impl *LanguagePackImpl) ModDescription() definition.ModDescription {
	return impl.modDescription
}

func (impl *LanguagePackImpl) GetCharacterDescription(longID uint64) (has bool, description definition.CharacterDescription) {
	if impl.characterDescription == nil {
		impl.characterDescription = map[uint64]definition.CharacterDescription{}
	}

	desc, exist := impl.characterDescription[longID]
	return exist, desc
}

func (impl *LanguagePackImpl) GetSkillDescription(longID uint64) (has bool, description definition.SkillDescription) {
	if impl.skillDescription == nil {
		impl.skillDescription = map[uint64]definition.SkillDescription{}
	}

	desc, exist := impl.skillDescription[longID]
	return exist, desc
}

func (impl *LanguagePackImpl) GetEventDescription(longID uint64) (has bool, description definition.EventDescription) {
	if impl.eventDescription == nil {
		impl.eventDescription = map[uint64]definition.EventDescription{}
	}

	desc, exist := impl.eventDescription[longID]
	return exist, desc
}

func (impl *LanguagePackImpl) GetCardDescription(longID uint64) (has bool, description definition.CardDescription) {
	if impl.cardDescription == nil {
		impl.cardDescription = map[uint64]definition.CardDescription{}
	}

	desc, exist := impl.cardDescription[longID]
	return exist, desc
}

func (impl *LanguagePackImpl) GetModifierDescription(longID uint64) (has bool, description definition.ModifierDescription) {
	if impl.modifierDescription == nil {
		impl.modifierDescription = map[uint64]definition.ModifierDescription{}
	}

	desc, exist := impl.modifierDescription[longID]
	return exist, desc
}

func (impl *LanguagePackImpl) AddCharacterDescription(description definition.CharacterDescription) {
	if impl.characterDescription == nil {
		impl.characterDescription = map[uint64]definition.CharacterDescription{}
	}

	impl.characterDescription[description.LongID()] = description
}

func (impl *LanguagePackImpl) AddSkillDescription(description definition.SkillDescription) {
	if impl.skillDescription == nil {
		impl.skillDescription = map[uint64]definition.SkillDescription{}
	}

	impl.skillDescription[description.LongID()] = description
}

func (impl *LanguagePackImpl) AddEventDescription(description definition.EventDescription) {
	if impl.eventDescription == nil {
		impl.eventDescription = map[uint64]definition.EventDescription{}
	}

	impl.eventDescription[description.LongID()] = description
}

func (impl *LanguagePackImpl) AddCardDescription(description definition.CardDescription) {
	if impl.cardDescription == nil {
		impl.cardDescription = map[uint64]definition.CardDescription{}
	}

	impl.cardDescription[description.LongID()] = description
}

func (impl *LanguagePackImpl) AddModifierDescription(description definition.ModifierDescription) {
	if impl.modifierDescription == nil {
		impl.modifierDescription = map[uint64]definition.ModifierDescription{}
	}

	impl.modifierDescription[description.LongID()] = description
}

type LanguagePackOptions func(option *LanguagePackImpl)

func WithLanguagePackLanguage(language enum.Language) LanguagePackOptions {
	return func(option *LanguagePackImpl) {
		option.language = language
	}
}

func WithLanguagePackModDescription(description definition.ModDescription) LanguagePackOptions {
	return func(option *LanguagePackImpl) {
		option.modDescription = description
	}
}

func NewLanguagePackWithOpts(options ...LanguagePackOptions) definition.LanguagePack {
	impl := &LanguagePackImpl{
		characterDescription: map[uint64]definition.CharacterDescription{},
		skillDescription:     map[uint64]definition.SkillDescription{},
		eventDescription:     map[uint64]definition.EventDescription{},
		cardDescription:      map[uint64]definition.CardDescription{},
		modifierDescription:  map[uint64]definition.ModifierDescription{},
	}

	for _, option := range options {
		option(impl)
	}

	return impl
}
