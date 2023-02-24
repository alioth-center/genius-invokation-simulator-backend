package testing

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
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

func RunCharacterImplements(mod definition.Mod, t *testing.T) {
	checkFunction := func(character definition.Character, t *testing.T) {
		defer protectRunFunction(t)
		character.Affiliation()
		character.Vision()
		character.Weapon()
		character.HP()
		character.MP()

		if character.TypeID() == 0 {
			t.Errorf("unassigned id for character %+v", character)
		}
		if character.Skills() == nil || len(character.Skills()) == 0 {
			t.Errorf("empty character skills in %v: %+v", character.TypeID(), character)
		}
		if character.Name() == "" {
			t.Errorf("empty character name in %v: %+v", character.TypeID(), character)
		}
	}

	type testCase struct {
		name      string
		character definition.Character
		checkFunc func(character definition.Character, t *testing.T)
	}

	t.Run("RunCharacterImplements", func(t *testing.T) {
		characters := mod.ProduceCharacters()
		tests := make([]testCase, len(characters))
		for i, character := range characters {
			tests[i] = testCase{
				name:      fmt.Sprintf("RunCharacterImplement-%v(%v)", character.Name(), character.TypeID()),
				character: character,
				checkFunc: checkFunction,
			}
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.checkFunc(tt.character, t)
			})
		}
	})
}
