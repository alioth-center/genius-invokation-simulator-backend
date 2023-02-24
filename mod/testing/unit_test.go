package testing

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"testing"
)

func TestCheckModImplementTestingFunction(t *testing.T) {
	mod := implement.NewMod()

	// 注释掉下一行测试结果则为fatal
	mod.RegisterCharacter(implement.NewCharacterWithOpts())

	t.Run("TestCheckModImplement", CheckModImplementTestingFunction(mod))
}

func TestCheckRepeatEntityIDTestingFunction(t *testing.T) {
	mod := implement.NewMod()

	mod.RegisterCharacter(implement.NewCharacterWithOpts(implement.WithCharacterID(1)))

	// 取消注释下一行测试结果则为fatal
	//mod.RegisterCard(implement.NewCardWithOpts(implement.WithCardID(1)))

	t.Run("TestCheckRepeatEntityID", CheckRepeatEntityIDTestingFunction(mod))
}

func TestRunCharacterImplementsTestingFunction(t *testing.T) {
	mod := implement.NewMod()

	mod.RegisterCharacter(implement.NewCharacterWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithCharacterID(114),
		implement.WithCharacterName("test"),
		implement.WithCharacterSkills(
			implement.NewSkillWithOpts(implement.WithSkillID(123)),
		),
	))

	t.Run("TestRunCharacterImplements", RunCharacterImplementsTestingFunction(mod))
}

func TestRunSkillImplementsTestingFunction(t *testing.T) {
	mod := implement.NewMod()
	mod.RegisterSkill(implement.NewAttackSkillWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithAttackSkillID(114),
		implement.WithAttackSkillCost(map[enum.ElementType]uint{}),
		implement.WithAttackSkillType(enum.SkillNormalAttack),
		implement.WithAttackSkillActiveDamageHandler(func(ctx definition.Context) (elementType enum.ElementType, damageAmount uint) {
			return enum.ElementUndefined, 0
		}),
		implement.WithAttackSkillBackgroundDamageHandler(func(ctx definition.Context) (damageAmount uint) {
			return 0
		}),
	))

	t.Run("TestRunSkillImplements", RunSkillImplementsTestingFunction(mod))
}

func TestRunEventImplementsTestingFunction(t *testing.T) {
	mod := implement.NewMod()
	mod.RegisterEvent(implement.NewEventWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithEventID(114),
		implement.WithEventTriggerAt(enum.AfterDefence),
		implement.WithEventCallback(func(ctx *context.CallbackContext) {}),
		implement.WithEventClearNow(func() bool { return true }),
		implement.WithEventTriggerNow(func(ctx context.CallbackContext) bool { return true }),
	))

	t.Run("TestRunEventImplements", RunEventImplementsTestingFunction(mod))
}

func TestRunCardImplementsTestingFunction(t *testing.T) {
	mod := implement.NewMod()
	mod.RegisterCard(implement.NewArtifactCardWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithArtifactCardID(114),
		implement.WithArtifactCardCost(map[enum.ElementType]uint{}),
		implement.WithArtifactCardModify(implement.NewEventWithOpts()),
	))

	t.Run("TestRunCardImplements", RunCardImplementsTestingFunction(mod))
}

type victorCalculator struct{}

func (impl victorCalculator) CalculateVictors(players []model.Player) (has bool, victors []model.Player) {
	return
}

type reactionCalculator struct{}

func (impl reactionCalculator) ReactionCalculate(types []enum.ElementType) (reaction enum.Reaction, elementRemains []enum.ElementType) {
	return
}

func (impl reactionCalculator) DamageCalculate(reaction enum.Reaction, targetCharacter model.Character, ctx *context.DamageContext) {
	return
}

func (impl reactionCalculator) EffectCalculate(reaction enum.Reaction, targetPlayer model.Player) (ctx *context.CallbackContext) {
	return
}

func (impl reactionCalculator) Attach(originalElements []enum.ElementType, newElement enum.ElementType) (resultElements []enum.ElementType) {
	return
}

func (impl reactionCalculator) Relative(reaction enum.Reaction, relativeElement enum.ElementType) bool {
	return true
}

func TestRunRuleImplementsTestingFunction(t *testing.T) {
	mod := implement.NewMod()
	mod.RegisterRule(implement.NewRuleWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithRuleID(114),
		implement.WithRuleImplement(enum.RuleTypeVictorCalculator, victorCalculator{}),
		implement.WithRuleImplement(enum.RuleTypeReactionCalculator, reactionCalculator{}),
	))

	t.Run("TestRunRuleImplements", RunRuleImplementsTestingFunction(mod))
}
