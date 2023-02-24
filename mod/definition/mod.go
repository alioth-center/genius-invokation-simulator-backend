package definition

type Mod interface {
	CharacterFactory
	SkillFactory
	EventFactory
	SummonFactory
	CardFactory
	RuleFactory
	LanguagePackFactory
	RegisterCharacter(character Character)
	RegisterSkill(skill Skill)
	RegisterEvent(event Event)
	RegisterSummon(summon Summon)
	RegisterCard(card Card)
	RegisterRule(rule Rule)
	AttachLanguagePack(languagePack LanguagePack)
}
