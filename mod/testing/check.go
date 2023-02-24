package testing

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"testing"
)

func CheckModImplement(mod definition.Mod, t *testing.T) {
	t.Run("CheckModImplement", func(t *testing.T) {
		totalImplementations := 0
		totalImplementations += len(mod.ProduceCharacters())
		totalImplementations += len(mod.ProduceSkill())
		totalImplementations += len(mod.ProduceEvents())
		totalImplementations += len(mod.ProduceSummons())
		totalImplementations += len(mod.ProduceCard())
		totalImplementations += len(mod.ProduceRule())

		if totalImplementations == 0 {
			t.Errorf("at least 1 implementation in a mod should be implemented")
		}
	})
}

func CheckRepeatEntityID(mod definition.Mod, t *testing.T) {
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
				skills := mod.ProduceSkill()
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
				cards := mod.ProduceCard()
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
				rules := mod.ProduceRule()
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
