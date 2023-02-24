package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type ModImpl struct {
	characters map[uint64]definition.Character
	skills     map[uint64]definition.Skill
	events     map[uint64]definition.Event
	summons    map[uint64]definition.Summon
	cards      map[uint64]definition.Card
	rules      map[uint64]definition.Rule
	languages  map[enum.Language]definition.LanguagePack
}

func (impl *ModImpl) ProduceCharacters() []definition.Character {
	if impl.characters == nil {
		impl.characters = map[uint64]definition.Character{}
	}

	result := make([]definition.Character, 0, len(impl.characters))
	for _, character := range impl.characters {
		result = append(result, character)
	}

	return result
}

func (impl *ModImpl) ProduceSkills() []definition.Skill {
	if impl.skills == nil {
		impl.skills = map[uint64]definition.Skill{}
	}

	result := make([]definition.Skill, 0, len(impl.skills))
	for _, skill := range impl.skills {
		result = append(result, skill)
	}

	return result
}

func (impl *ModImpl) ProduceEvents() []definition.Event {
	if impl.events == nil {
		impl.events = map[uint64]definition.Event{}
	}

	result := make([]definition.Event, 0, len(impl.events))
	for _, event := range impl.events {
		result = append(result, event)
	}

	return result
}

func (impl *ModImpl) ProduceSummons() []definition.Summon {
	if impl.summons == nil {
		impl.summons = map[uint64]definition.Summon{}
	}

	result := make([]definition.Summon, 0, len(impl.summons))
	for _, summon := range impl.summons {
		result = append(result, summon)
	}

	return result
}

func (impl *ModImpl) ProduceCards() []definition.Card {
	if impl.cards == nil {
		impl.cards = map[uint64]definition.Card{}
	}

	result := make([]definition.Card, 0, len(impl.cards))
	for _, card := range impl.cards {
		result = append(result, card)
	}

	return result
}

func (impl *ModImpl) ProduceRules() []definition.Rule {
	if impl.rules == nil {
		impl.rules = map[uint64]definition.Rule{}
	}

	result := make([]definition.Rule, 0, len(impl.rules))
	for _, rule := range impl.rules {
		result = append(result, rule)
	}

	return result
}

func (impl *ModImpl) ProduceLanguagePacks() []definition.LanguagePack {
	if impl.languages == nil {
		impl.languages = map[enum.Language]definition.LanguagePack{}
	}

	result := make([]definition.LanguagePack, 0, len(impl.languages))
	for _, languagePack := range impl.languages {
		result = append(result, languagePack)
	}

	return result
}

func (impl *ModImpl) RegisterCharacter(character definition.Character) {
	if impl.characters == nil {
		impl.characters = map[uint64]definition.Character{}
	}

	impl.characters[character.TypeID()] = character
}

func (impl *ModImpl) RegisterSkill(skill definition.Skill) {
	if impl.skills == nil {
		impl.skills = map[uint64]definition.Skill{}
	}

	impl.skills[skill.TypeID()] = skill
}

func (impl *ModImpl) RegisterEvent(event definition.Event) {
	if impl.events == nil {
		impl.events = map[uint64]definition.Event{}
	}

	impl.events[event.TypeID()] = event
}

func (impl *ModImpl) RegisterSummon(summon definition.Summon) {
	if impl.summons == nil {
		impl.summons = map[uint64]definition.Summon{}
	}

	impl.summons[summon.TypeID()] = summon
}

func (impl *ModImpl) RegisterCard(card definition.Card) {
	impl.cards[card.TypeID()] = card
}

func (impl *ModImpl) RegisterRule(rule definition.Rule) {
	impl.rules[rule.TypeID()] = rule
}

func (impl *ModImpl) AttachLanguagePack(languagePack definition.LanguagePack) {
	impl.languages[languagePack.Language()] = languagePack
}

func NewMod() definition.Mod {
	return &ModImpl{
		characters: map[uint64]definition.Character{},
		skills:     map[uint64]definition.Skill{},
		events:     map[uint64]definition.Event{},
		summons:    map[uint64]definition.Summon{},
		cards:      map[uint64]definition.Card{},
		rules:      map[uint64]definition.Rule{},
		languages:  map[enum.Language]definition.LanguagePack{},
	}
}
