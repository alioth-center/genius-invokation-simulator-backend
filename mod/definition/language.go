package definition

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type LanguagePack interface {
	Language() enum.Language
	ModDescription() ModDescription
	GetCharacterDescription(longID uint64) (has bool, description CharacterDescription)
	GetSkillDescription(longID uint64) (has bool, description SkillDescription)
	GetEventDescription(longID uint64) (has bool, description EventDescription)
	GetCardDescription(longID uint64) (has bool, description CardDescription)
	GetSummonDescription(longID uint64) (has bool, description SummonDescription)
	GetModifierDescription(longID uint64) (has bool, description ModifierDescription)
	AddCharacterDescription(description CharacterDescription)
	AddSkillDescription(description SkillDescription)
	AddEventDescription(description EventDescription)
	AddCardDescription(description CardDescription)
	AddModifierDescription(description ModifierDescription)
}
