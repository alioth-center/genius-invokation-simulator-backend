package definition

type Mod interface {
	CharacterFactory
	SkillFactory
	EventFactory
	SummonFactory
	CardFactory
	RuleFactory
	RegisterCharacter(character Character)
	RegisterSkill(skill Skill)
	RegisterEvent(event Event)
	RegisterSummon(summon Summon)
	RegisterCard(card Card)
	RegisterRule(rule Rule)
}
