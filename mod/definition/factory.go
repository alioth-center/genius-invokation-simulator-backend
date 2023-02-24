package definition

type CharacterFactory interface {
	ProduceCharacters() []Character
}

type SummonFactory interface {
	ProduceSummons() []Summon
}

type EventFactory interface {
	ProduceEvents() []Event
}

type SkillFactory interface {
	ProduceSkills() []Skill
}

type CardFactory interface {
	ProduceCards() []Card
}

type RuleFactory interface {
	ProduceRules() []Rule
}

type LanguagePackFactory interface {
	ProduceLanguagePacks() []LanguagePack
}
