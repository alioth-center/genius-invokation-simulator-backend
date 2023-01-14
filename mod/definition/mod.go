package adapter

type Mod interface {
	CharacterFactory
	SkillFactory
	EventFactory
	SummonFactory
	CardFactory
	RuleFactory
}
