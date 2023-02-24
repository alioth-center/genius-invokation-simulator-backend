package definition

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type Description interface {
	DescriptionType() enum.DescriptionType
	ShortID() uint16
	LongID() uint64
}

type CharacterDescription interface {
	Description
	CharacterName() string
	CharacterDescription() string
	CharacterTitle() string
	CharacterStory() string
}

type SkillDescription interface {
	Description
	SkillName() string
	SkillDescription() string
}

type EventDescription interface {
	Description
	EventName() string
	EventDescription() string
}

type CardDescription interface {
	Description
	CardName() string
	CardDescription() string
}

type SummonDescription interface {
	Description
	SummonName() string
	SummonDescription() string
}

type ModifierDescription interface {
	Description
	ModifierName() string
	ModifierDescription() string
}

type ModDescription interface {
	ModID() uint64
	ModName() string
	ModDescription() string
}
