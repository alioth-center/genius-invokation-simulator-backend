/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "UndividedHeart.go" LastUpdatedAt 2022/12/14 08:46:14
 */

package talent

import "github.com/sunist-c/genius-invokation-simulator-backend/model"

type UndividedHeart struct {
	triggerSkill model.ISkill
	activated    bool
}

func (u *UndividedHeart) DamageModifier() model.AttackDamageModifier {
	return func(ctx *model.ModifierContext[model.AttackDamageContext]) {
		if u.triggerSkill.Name() == ctx.Data.Sender.Name() {
			if u.activated {
				ctx.Data.AddActiveDamage(1)
				ctx.Data.AddPenetratedDamage(1)
			} else {
				u.activated = true
			}
		}

		ctx.Continue()
	}
}
