package testing

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"runtime"
	"testing"
)

func protectRunFunction(t *testing.T) {
	err := recover()

	if err != nil {
		switch err.(type) {
		case runtime.Error:
			t.Errorf("runtime error: %v", err)
		default:
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func RunCharacterImplementsTestingFunction(mod definition.Mod) func(t *testing.T) {
	checkFunction := func(character definition.Character, t *testing.T) {
		defer protectRunFunction(t)

		if character.TypeID() == 0 {
			t.Errorf("unassigned id for character %+v", character)
		}
		if character.Skills() == nil || len(character.Skills()) == 0 {
			t.Errorf("empty character skills in %v: %+v", character.TypeID(), character)
		}
		if character.Name() == "" {
			t.Errorf("empty character name in %v: %+v", character.TypeID(), character)
		}

		character.Affiliation()
		character.Vision()
		character.Weapon()
		character.HP()
		character.MP()
	}

	type testCase struct {
		name      string
		character definition.Character
	}

	return func(t *testing.T) {
		characters := mod.ProduceCharacters()
		tests := make([]testCase, len(characters))
		for i, character := range characters {
			tests[i] = testCase{
				name:      fmt.Sprintf("RunCharacterImplement-%v(%v)", character.Name(), character.TypeID()),
				character: character,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkFunction(tt.character, t)
			})
		}
	}
}

func RunSkillImplementsTestingFunction(mod definition.Mod) func(t *testing.T) {
	checkFunction := func(skill definition.Skill, t *testing.T) {
		defer protectRunFunction(t)

		if skill.TypeID() == 0 {
			t.Errorf("unassigned id for skill %+v", skill)
		}

		skill.SkillType()

		if attackSkill, success := skill.(definition.AttackSkill); success {
			if attackSkill.SkillCost() == nil {
				t.Errorf("nil skill cost in attack skill %v: %+v", attackSkill.TypeID(), attackSkill)
			}
			if attackSkill.ActiveDamage(implement.NewEmptyContext()) == nil {
				t.Errorf("nil active damage in attack skill %v: %+v", attackSkill.TypeID(), attackSkill)
			}
			if attackSkill.BackgroundDamage(implement.NewEmptyContext()) == nil {
				t.Errorf("nil background damage in attack skill %v: %+v", attackSkill.TypeID(), attackSkill)
			}
		}
	}

	type testCase struct {
		name  string
		skill definition.Skill
	}

	return func(t *testing.T) {
		skills := mod.ProduceSkill()
		tests := make([]testCase, len(skills))
		for i, skill := range skills {
			tests[i] = testCase{
				name:  fmt.Sprintf("RunSkillImplement-%v", skill.TypeID()),
				skill: skill,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkFunction(tt.skill, t)
			})
		}
	}
}

func RunEventImplementsTestingFunction(mod definition.Mod) func(t *testing.T) {
	checkFunction := func(event definition.Event, t *testing.T) {
		defer protectRunFunction(t)

		if event.TypeID() == 0 {
			t.Errorf("unassigned id for event %+v", event)
		}

		event.TriggerAt()

		ctx := context.NewCallbackContext()
		event.TriggeredNow(*ctx)
		event.CallBack(ctx)
		event.ClearNow()
	}

	type testCase struct {
		name  string
		event definition.Event
	}

	return func(t *testing.T) {
		events := mod.ProduceEvents()
		tests := make([]testCase, len(events))
		for i, event := range events {
			tests[i] = testCase{
				name:  fmt.Sprintf("RunEventImplement-%v", event.TypeID()),
				event: event,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkFunction(tt.event, t)
			})
		}
	}
}

func RunCardImplementsTestingFunction(mod definition.Mod) func(t *testing.T) {
	checkFunction := func(card definition.Card, t *testing.T) {
		defer protectRunFunction(t)

		if card.TypeID() == 0 {
			t.Errorf("unassigned id for card %+v", card)
		}
		if card.Cost() == nil {
			t.Errorf("nil cost in card %v: %+v", card.TypeID(), card)
		}

		card.CardType()

		if eventCard, success := card.(definition.EventCard); success {
			if eventCard.Event() == nil {
				t.Errorf("nil event in event card %v: %+v", eventCard.TypeID(), eventCard)
			}
		}

		if supportCard, success := card.(definition.SupportCard); success {
			if supportCard.Support() == nil {
				t.Errorf("nil support in support card %v: %+v", supportCard.TypeID(), supportCard)
			}
		}

		if equipmentCard, success := card.(definition.EquipmentCard); success {
			if equipmentCard.Modify() == nil {
				t.Errorf("nil equipment modifier in equipment card %v: %+v", equipmentCard.TypeID(), equipmentCard)
			}

			equipmentCard.EquipmentType()

			if weaponCard, success := card.(definition.WeaponCard); success {
				weaponCard.WeaponType()
			}
		}
	}

	type testCase struct {
		name string
		card definition.Card
	}

	return func(t *testing.T) {
		cards := mod.ProduceCard()
		tests := make([]testCase, len(cards))
		for i, card := range cards {
			tests[i] = testCase{
				name: fmt.Sprintf("RunCardImplement-%v", card.TypeID()),
				card: card,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkFunction(tt.card, t)
			})
		}
	}
}

func RunRuleImplementsTestingFunction(mod definition.Mod) func(t *testing.T) {
	checkFunction := func(rule definition.Rule, t *testing.T) {
		defer protectRunFunction(t)

		if rule.TypeID() == 0 {
			t.Errorf("unassigned id for rule %+v", rule)
		}
		if !rule.CheckImplements() {
			t.Errorf("check failed in rule %v: %+v", rule.TypeID(), rule)
		}

		if victorCalculator := rule.Implements(enum.RuleTypeVictorCalculator); victorCalculator == nil {
			t.Errorf("victor calculator is not exist in rule %v: %+v", rule.TypeID(), rule)
		} else if realVictorCalculator, success := victorCalculator.(model.VictorCalculator); !success {
			t.Errorf("victor calculator is not implemented in rule %v: %+v", rule.TypeID(), rule)
		} else {
			// todo: complete test cases
			realVictorCalculator.CalculateVictors([]model.Player{})
		}

		if reactionCalculator := rule.Implements(enum.RuleTypeReactionCalculator); reactionCalculator == nil {
			t.Errorf("reaction calculator is not exist in rule %v: %+v", rule.TypeID(), rule)
		} else if realReactionCalculator, success := reactionCalculator.(model.ReactionCalculator); !success {
			t.Errorf("reaction calculator is not implemented in rule %v: %+v", rule.TypeID(), rule)
		} else {
			// todo: complete test cases
			realReactionCalculator.ReactionCalculate([]enum.ElementType{})
			realReactionCalculator.EffectCalculate(enum.ReactionNone, nil)
			realReactionCalculator.Attach([]enum.ElementType{}, enum.ElementNone)
			realReactionCalculator.DamageCalculate(enum.ReactionNone, nil, nil)
			realReactionCalculator.Relative(enum.ReactionNone, enum.ElementNone)
		}
	}

	type testCase struct {
		name string
		rule definition.Rule
	}

	return func(t *testing.T) {
		rules := mod.ProduceRule()
		tests := make([]testCase, len(rules))
		for i, rule := range rules {
			tests[i] = testCase{
				name: fmt.Sprintf("RunRuleImplement-%v", rule.TypeID()),
				rule: rule,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				checkFunction(tt.rule, t)
			})
		}
	}
}
