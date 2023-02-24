package testing

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"testing"
)

func TestCheckModImplement(t *testing.T) {
	mod := implement.NewMod()

	// 注释掉下一行测试结果则为fatal
	mod.RegisterCharacter(implement.NewCharacterWithOpts())

	t.Run("TestCheckModImplement", func(t *testing.T) {
		CheckModImplement(mod, t)
	})
}

func TestCheckRepeatEntityID(t *testing.T) {
	mod := implement.NewMod()

	mod.RegisterCharacter(implement.NewCharacterWithOpts(implement.WithCharacterID(1)))

	// 取消注释下一行测试结果则为fatal
	// mod.RegisterCard(implement.NewCardWithOpts(implement.WithCardID(1)))

	t.Run("TestCheckRepeatEntityID", func(t *testing.T) {
		CheckRepeatEntityID(mod, t)
	})
}

func TestRunCharacterImplements(t *testing.T) {
	mod := implement.NewMod()

	mod.RegisterCharacter(implement.NewCharacterWithOpts(
		// 注释掉Options测试结果则为fatal
		implement.WithCharacterID(114),
		implement.WithCharacterName("test"),
		implement.WithCharacterSkills(
			implement.NewSkillWithOpts(implement.WithSkillID(123)),
		),
	))

	t.Run("TestRunCharacterImplements", func(t *testing.T) {
		RunCharacterImplements(mod, t)
	})
}
