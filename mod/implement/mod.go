package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
)

type ModImpl struct {
	characters map[uint64]definition.Character
	skills     map[uint64]definition.Skill
	events     map[uint64]definition.Event
	summons    map[uint64]definition.Summon
	cards      map[uint64]definition.Card
	rules      map[uint64]definition.Rule
}

func (impl *ModImpl) ProduceCharacters() []definition.Character {
	if impl.characters == nil {
		impl.characters = map[uint64]definition.Character{}
	}

	result := make([]definition.Character, len(impl.characters), 0)
	for _, character := range impl.characters {
		result = append(result, character)
	}

	return result
}

func (impl *ModImpl) ProduceSkill() []definition.Skill {
	if impl.skills == nil {
		impl.skills = map[uint64]definition.Skill{}
	}

	result := make([]definition.Skill, len(impl.skills), 0)
	for _, skill := range impl.skills {
		result = append(result, skill)
	}

	return result
}

func (impl *ModImpl) ProduceEvents() []definition.Event {
	if impl.events == nil {
		impl.events = map[uint64]definition.Event{}
	}

	result := make([]definition.Event, len(impl.events), 0)
	for _, event := range impl.events {
		result = append(result, event)
	}

	return result
}

func (impl *ModImpl) ProduceSummons() []definition.Summon {
	if impl.summons == nil {
		impl.summons = map[uint64]definition.Summon{}
	}

	result := make([]definition.Summon, len(impl.summons), 0)
	for _, summon := range impl.summons {
		result = append(result, summon)
	}

	return result
}

func (impl *ModImpl) ProduceCard() []definition.Card {
	if impl.cards == nil {
		impl.cards = map[uint64]definition.Card{}
	}

	result := make([]definition.Card, len(impl.cards), 0)
	for _, card := range impl.cards {
		result = append(result, card)
	}

	return result
}

func (impl *ModImpl) ProduceRule() []definition.Rule {
	if impl.rules == nil {
		impl.rules = map[uint64]definition.Rule{}
	}

	result := make([]definition.Rule, len(impl.rules), 0)
	for _, rule := range impl.rules {
		result = append(result, rule)
	}

	return result
}

func (impl *ModImpl) RegisterCharacter(character definition.Character) {
	impl.characters[character.TypeID()] = character
}

func (impl *ModImpl) RegisterSkill(skill definition.Skill) {
	impl.skills[skill.TypeID()] = skill
}

func (impl *ModImpl) RegisterEvent(event definition.Event) {
	impl.events[event.TypeID()] = event
}

func (impl *ModImpl) RegisterSummon(summon definition.Summon) {
	impl.summons[summon.TypeID()] = summon
}

func (impl *ModImpl) RegisterCard(card definition.Card) {
	impl.cards[card.TypeID()] = card
}

func (impl *ModImpl) RegisterRule(rule definition.Rule) {
	impl.rules[rule.TypeID()] = rule
}

func NewMod() definition.Mod {
	return &ModImpl{
		characters: map[uint64]definition.Character{},
		skills:     map[uint64]definition.Skill{},
		events:     map[uint64]definition.Event{},
		summons:    map[uint64]definition.Summon{},
		cards:      map[uint64]definition.Card{},
		rules:      map[uint64]definition.Rule{},
	}
}
