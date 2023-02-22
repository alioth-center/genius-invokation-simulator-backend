package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"testing"
)

func TestNextID(t *testing.T) {
	t.Run("TestNextID", func(t *testing.T) {
		for i := 0; i < 65536; i++ {
			if got := NextID(); got != uint16(i+1) {
				t.Errorf("unexpected next id: %d, want: %d", got, i+1)
			}
		}
	})

	usedID = map[uint16]bool{}
}

func TestUseID(t *testing.T) {
	t.Run("TestUseID", func(t *testing.T) {
		UseID(2)
		UseID(5)
		NextID()
		if got := NextID(); got != 3 {
			t.Errorf("unexpected use id: %d, got %v", 2, got)
		}
		NextID()
		if got := NextID(); got != 6 {
			t.Errorf("unexpected use id: %d got %v", 5, got)
		}
	})

	usedID = map[uint16]bool{}
}

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
				WithCharacterSkills(
					// 普通攻击
					NewAttackSkillWithOpts(
						WithAttackSkillID(101),
						WithAttackSkillType(enum.SkillNormalAttack),
						WithAttackSkillCost(map[enum.ElementType]uint{
							enum.ElementCryo:     1,
							enum.ElementCurrency: 2,
						}),
						WithAttackSkillActiveDamageHandler(func(ctx definition.Context) (elementType enum.ElementType, damageAmount uint) {
							return enum.ElementNone, 2
						}),
					),

					// 霜华矢
					NewAttackSkillWithOpts(
						WithAttackSkillID(101),
						WithAttackSkillType(enum.SkillNormalAttack),
						WithAttackSkillCost(map[enum.ElementType]uint{
							enum.ElementCryo: 5,
						}),
						WithAttackSkillActiveDamageHandler(func(ctx definition.Context) (elementType enum.ElementType, damageAmount uint) {
							return enum.ElementCryo, 2
						}),
						WithAttackSkillBackgroundDamageHandler(func(ctx definition.Context) (damageAmount uint) {
							return 2
						}),
					),
				),
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

func TestNewRuleWithOpts(t *testing.T) {
	SetDebugFlag(true)

	var nilReactionCalculatorImpl model.ReactionCalculator
	var nilVictorCalculatorImpl model.VictorCalculator
	var defaultImpl = NewRuleWithOpts(WithRuleID(1), WithRuleImplement(enum.RuleTypeVictorCalculator, nilVictorCalculatorImpl))

	tests := []struct {
		name    string
		options []RuleImplOptions
		want    func(rule definition.Rule) bool
	}{
		{
			name: "TestNewRuleWithOpts-Common-Success",
			options: []RuleImplOptions{
				WithRuleID(1),
				WithRuleImplement(enum.RuleTypeReactionCalculator, nilReactionCalculatorImpl),
				WithRuleImplement(enum.RuleTypeVictorCalculator, nilVictorCalculatorImpl),
			},
			want: func(rule definition.Rule) bool {
				entityImpl := NewEntityWithOpts(WithEntityID(1))
				if rule.TypeID() != entityImpl.TypeID() {
					return false
				}

				if rule.Implements(enum.RuleTypeReactionCalculator) != nilReactionCalculatorImpl {
					return false
				}

				if rule.Implements(enum.RuleTypeVictorCalculator) != nilVictorCalculatorImpl {
					return false
				}

				return true
			},
		},
		{
			name: "TestNewRuleWithOpts-Common-CheckFailed",
			options: []RuleImplOptions{
				WithRuleID(1),
			},
			want: func(rule definition.Rule) bool {
				entityImpl := NewEntityWithOpts(WithEntityID(1))
				if rule.TypeID() != entityImpl.TypeID() {
					return false
				}

				if rule.Implements(enum.RuleTypeReactionCalculator) != nil {
					return false
				}

				if rule.Implements(enum.RuleTypeVictorCalculator) != nil {
					return false
				}

				return true
			},
		},
		{
			name: "TestNewRuleWithOpts-CopyFrom-Success",
			options: []RuleImplOptions{
				WithRuleID(1),
				WithRuleImplement(enum.RuleTypeReactionCalculator, nilReactionCalculatorImpl),
				WithRuleCopyFrom(defaultImpl, enum.RuleTypeVictorCalculator),
			},
			want: func(rule definition.Rule) bool {
				entityImpl := NewEntityWithOpts(WithEntityID(1))
				if rule.TypeID() != entityImpl.TypeID() {
					return false
				}

				if rule.Implements(enum.RuleTypeVictorCalculator) != nilVictorCalculatorImpl {
					return false
				}

				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := NewRuleWithOpts(tt.options...)
			if !tt.want(impl) {
				t.Errorf("unexpected rule impl: %+v", impl)
			}
		})
	}
}
