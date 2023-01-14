package adapter

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
	ProduceSkill() []Skill
}

type CardFactory interface {
	ProduceCard() []Card
}

type RuleFactory interface {
	ProduceRuleSet() []Rule
}
