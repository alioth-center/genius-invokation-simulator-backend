package testing

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"testing"
)

func CheckModImplementTestingFunction(mod definition.Mod) func(t *testing.T) {
	return func(t *testing.T) {
		totalImplementations := 0
		totalImplementations += len(mod.ProduceCharacters())
		totalImplementations += len(mod.ProduceSkills())
		totalImplementations += len(mod.ProduceEvents())
		totalImplementations += len(mod.ProduceSummons())
		totalImplementations += len(mod.ProduceCards())
		totalImplementations += len(mod.ProduceRules())

		if totalImplementations == 0 {
			t.Errorf("at least 1 implementation in a mod should be implemented")
		}
	}
}

func CheckRepeatEntityIDTestingFunction(mod definition.Mod) func(t *testing.T) {
	return func(t *testing.T) {
		checkList := map[uint64]struct{}{}
		tests := []struct {
			name      string
			checkFunc func(mod definition.Mod, t *testing.T)
		}{
			{
				name: "CheckRepeatCharacterID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					characters := mod.ProduceCharacters()
					for _, character := range characters {
						if _, exist := checkList[character.TypeID()]; exist {
							t.Errorf("id %v for character %+v is already in use", character.TypeID(), character)
						} else {
							checkList[character.TypeID()] = struct{}{}
						}
					}
				},
			},
			{
				name: "CheckRepeatSkillID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					skills := mod.ProduceSkills()
					for _, skill := range skills {
						if _, exist := checkList[skill.TypeID()]; exist {
							t.Errorf("id %v for skill %+v is already in use", skill.TypeID(), skill)
						} else {
							checkList[skill.TypeID()] = struct{}{}
						}
					}
				},
			},
			{
				name: "CheckRepeatEventID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					events := mod.ProduceEvents()
					for _, event := range events {
						if _, exist := checkList[event.TypeID()]; exist {
							t.Errorf("id %v for event %+v is already in use", event.TypeID(), event)
						} else {
							checkList[event.TypeID()] = struct{}{}
						}
					}
				},
			},
			{
				name: "CheckRepeatSummonID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					summons := mod.ProduceSummons()
					for _, summon := range summons {
						if _, exist := checkList[summon.TypeID()]; exist {
							t.Errorf("id %v for summon %+v is already in use", summon.TypeID(), summon)
						} else {
							checkList[summon.TypeID()] = struct{}{}
						}
					}
				},
			},
			{
				name: "CheckRepeatCardID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					cards := mod.ProduceCards()
					for _, card := range cards {
						if _, exist := checkList[card.TypeID()]; exist {
							t.Errorf("id %v for card %+v is already in use", card.TypeID(), card)
						} else {
							checkList[card.TypeID()] = struct{}{}
						}
					}
				},
			},
			{
				name: "CheckRepeatRuleID",
				checkFunc: func(mod definition.Mod, t *testing.T) {
					rules := mod.ProduceRules()
					for _, rule := range rules {
						if _, exist := checkList[rule.TypeID()]; exist {
							t.Errorf("id %v for rule %+v is already in use", rule.TypeID(), rule)
						} else {
							checkList[rule.TypeID()] = struct{}{}
						}
					}
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.checkFunc(mod, t)
			})
		}
	}
}

func CheckDescriptionsTestingFunction(mod definition.Mod) func(t *testing.T) {
	return func(t *testing.T) {
		languages := mod.ProduceLanguagePacks()
		if len(languages) == 0 {
			t.Errorf("at least one language in a mod should be attached to")
		}

		hasDescription := map[uint64]struct{}{}
		characters, skills, events, cards, summons := mod.ProduceCharacters(), mod.ProduceSkills(), mod.ProduceEvents(), mod.ProduceCards(), mod.ProduceSummons()
		for _, language := range languages {
			for _, character := range characters {
				if has, _ := language.GetCharacterDescription(character.TypeID()); !has {
					t.Logf("character %v not has a description in %v", character.TypeID(), implement.LanguageEnumToString(language.Language()))
				} else {
					hasDescription[character.TypeID()] = struct{}{}
				}
			}
			for _, skill := range skills {
				if has, _ := language.GetSkillDescription(skill.TypeID()); !has {
					t.Logf("skill %v not has a description in %v", skill.TypeID(), implement.LanguageEnumToString(language.Language()))
				} else {
					hasDescription[skill.TypeID()] = struct{}{}
				}
			}
			for _, event := range events {
				if has, _ := language.GetEventDescription(event.TypeID()); !has {
					t.Logf("event %v not has a description in %v", event.TypeID(), implement.LanguageEnumToString(language.Language()))
				} else {
					hasDescription[event.TypeID()] = struct{}{}
				}
			}
			for _, card := range cards {
				if has, _ := language.GetCardDescription(card.TypeID()); !has {
					t.Logf("card %v not has a description in %v", card.TypeID(), implement.LanguageEnumToString(language.Language()))
				} else {
					hasDescription[card.TypeID()] = struct{}{}
				}
			}
			for _, summon := range summons {
				if has, _ := language.GetSummonDescription(summon.TypeID()); !has {
					t.Logf("summon %v not has a description in %v", summon.TypeID(), implement.LanguageEnumToString(language.Language()))
				} else {
					hasDescription[summon.TypeID()] = struct{}{}
				}
			}
		}

		for _, character := range characters {
			if _, has := hasDescription[character.TypeID()]; !has {
				t.Errorf("character %v not has at least one description", character.TypeID())
			}
		}
		for _, skill := range skills {
			if _, has := hasDescription[skill.TypeID()]; !has {
				t.Errorf("skill %v not has at least one description", skill.TypeID())
			}
		}
		for _, event := range events {
			if _, has := hasDescription[event.TypeID()]; !has {
				t.Errorf("event %v not has at least one description", event.TypeID())
			}
		}
		for _, card := range cards {
			if _, has := hasDescription[card.TypeID()]; !has {
				t.Errorf("card %v not has at least one description", card.TypeID())
			}
		}
		for _, summon := range summons {
			if _, has := hasDescription[summon.TypeID()]; !has {
				t.Errorf("summon %v not has at least one description", summon.TypeID())
			}
		}
		// todo: add modifier description check
	}
}
