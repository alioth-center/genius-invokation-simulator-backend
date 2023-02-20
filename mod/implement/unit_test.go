package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"testing"
)

func TestNewCharacterWithOpts(t *testing.T) {
	SetDebugFlag(true)

	tests := []struct {
		name    string
		options []CharacterOptions
		want    func(character definition.Character) bool
	}{
		{
			name: "TestNewCharacterWithOpts-Common",
			options: []CharacterOptions{
				WithCharacterID(1),
				WithCharacterName("Ganyu"),
				WithCharacterAffiliation(enum.AffiliationLiyue),
				WithCharacterHP(10),
				WithCharacterMP(2),
				WithCharacterVision(enum.ElementCryo),
				WithCharacterWeapon(enum.WeaponBow),
			},
			want: func(character definition.Character) bool {
				entity := NewEntityWithOpts(WithEntityID(1))
				if character.TypeID() != entity.TypeID() {
					return false
				} else if character.Name() != "Ganyu" {
					return false
				} else if character.Affiliation() != enum.AffiliationLiyue {
					return false
				} else if character.HP() != 10 {
					return false
				} else if character.MP() != 2 {
					return false
				} else if character.Vision() != enum.ElementCryo {
					return false
				} else if character.Weapon() != enum.WeaponBow {
					return false
				} else {
					return true
				}
			},
		},
		{
			name: "TestNewCharacterWithOpts-Overwrite",
			options: []CharacterOptions{
				WithCharacterID(1),
				WithCharacterName("Ganyu"),
				WithCharacterAffiliation(enum.AffiliationLiyue),
				WithCharacterHP(10),
				WithCharacterMP(2),
				WithCharacterVision(enum.ElementCryo),
				WithCharacterWeapon(enum.WeaponBow),
				WithCharacterID(2),
			},
			want: func(character definition.Character) bool {
				entity := NewEntityWithOpts(WithEntityID(2))
				if character.TypeID() != entity.TypeID() {
					return false
				} else if character.Name() != "Ganyu" {
					return false
				} else if character.Affiliation() != enum.AffiliationLiyue {
					return false
				} else if character.HP() != 10 {
					return false
				} else if character.MP() != 2 {
					return false
				} else if character.Vision() != enum.ElementCryo {
					return false
				} else if character.Weapon() != enum.WeaponBow {
					return false
				} else {
					return true
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			character := NewCharacterWithOpts(tt.options...)
			if !tt.want(character) {
				t.Errorf("unexpected character: %+v", character)
			}
		})
	}
}
