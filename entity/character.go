package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

var (
	nullDirectAttackModifiers kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullFinalAttackModifiers  kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullDefenceModifiers      kv.Map[uint, []modifier.Modifier[context.DamageContext]] = nil
	nullChargeModifiers       kv.Map[uint, []modifier.Modifier[context.ChargeContext]] = nil
	nullHealModifiers         kv.Map[uint, []modifier.Modifier[context.HealContext]]   = nil
	nullCostModifiers         kv.Map[uint, []modifier.Modifier[context.CostContext]]   = nil
)

type Character interface {
	// ID 角色的ID
	ID() uint

	// Name 角色的名字
	Name() string

	// Description 角色的描述
	Description() string

	// Affiliation 角色的归属地
	Affiliation() enum.Affiliation

	// Vision 角色的元素
	Vision() enum.ElementType

	// Weapon 角色的武器类型
	Weapon() enum.WeaponType

	// MaxHP 角色的最大HP
	MaxHP() uint

	// MaxMP 角色的最大MP
	MaxMP() uint

	// HP 角色的当前HP
	HP() uint

	// MP 角色的当前MP
	MP() uint

	// Status 角色的状态
	Status() enum.CharacterStatus

	// ExecuteCharge 根据ChargeContext给角色增加或减少MP
	ExecuteCharge(ctx *context.ChargeContext)

	// ExecuteHeal 根据HealContext给角色进行治疗
	ExecuteHeal(ctx *context.HealContext)

	// ExecuteCostModify 使用角色的CostModifiers对CostContext进行修正
	ExecuteCostModify(ctx *context.CostContext)

	// ExecuteModify 根据ModifierContext对角色的Modifiers进行修改
	ExecuteModify(ctx *context.ModifierContext)

	// ExecuteDefence 根据DamageContext对角色进行伤害结算，不包括效果结算
	ExecuteDefence(ctx *context.DamageContext)

	// ExecuteAttack 使用skill进行对target角色和background后台角色进行攻击
	ExecuteAttack(skill, target uint, background []uint) (ctx *context.DamageContext)

	// ExecuteDirectAttackModifiers 使用角色的DirectAttackModifiers对DamageContext进行伤害修正
	ExecuteDirectAttackModifiers(ctx *context.DamageContext)

	// ExecuteFinalAttackModifiers 使用角色的FinalAttackModifiers对DamageContext进行伤害修正
	ExecuteFinalAttackModifiers(ctx *context.DamageContext)
}

type character struct {
	id          uint                // id 角色的ID，由框架确定
	player      uint                // player 所属玩家的ID，由框架确定
	name        string              // name 角色的名称
	description string              // description 角色的描述
	affiliation enum.Affiliation    // affiliation 角色的势力归属
	vision      enum.ElementType    // vision 角色的元素类型
	weapon      enum.WeaponType     // weapon 角色的武器类型
	skills      kv.Map[uint, Skill] // skills 角色的技能

	maxHP      uint                      // maxHP 角色的最大生命值
	currentHP  uint                      // currentHP 角色的当前生命值
	maxMP      uint                      // maxMP 角色的最大能量值
	currentMP  uint                      // currentMP 角色的当前能量值
	status     enum.CharacterStatus      // status 角色的状态
	elements   []enum.ElementType        // elements 角色目前附着的元素
	satiety    bool                      // satiety 角色的饱腹状态
	equipments kv.Map[uint, interface{}] // equipments 角色穿着的装备

	localDirectAttackModifiers AttackModifiers  // localDirectAttackModifiers 本地直接攻击修正
	localFinalAttackModifiers  AttackModifiers  // localFinalAttackModifiers 本地最终攻击修正
	localDefenceModifiers      DefenceModifiers // localDefenceModifiers 本地防御修正
	localChargeModifiers       ChargeModifiers  // localChargeModifiers 本地充能修正
	localHealModifiers         HealModifiers    // localHealModifiers 本地治疗修正
	localCostModifiers         CostModifiers    // localCostModifiers 本地费用修正
}

func (c *character) ExecuteCharge(ctx *context.ChargeContext) {
	c.localChargeModifiers.Execute(ctx)
	executeAmount := ctx.Charge()[c.id]

	if executeAmount > 0 {
		if c.currentMP+uint(executeAmount) > c.maxMP {
			c.currentMP = c.maxMP
		} else {
			c.currentMP += uint(executeAmount)
		}
	} else {
		if c.currentMP-uint(executeAmount) < 0 {
			c.currentMP = 0
		} else {
			c.currentMP -= uint(executeAmount)
		}
	}
}

func (c *character) ExecuteHeal(ctx *context.HealContext) {
	c.localHealModifiers.Execute(ctx)
	executeAmount := ctx.Heal()[c.id]

	if c.currentHP+executeAmount > c.maxHP {
		c.currentHP = c.maxHP
	} else {
		c.currentHP += executeAmount
	}
}

func (c *character) ExecuteCostModify(ctx *context.CostContext) {
	c.localCostModifiers.Execute(ctx)
}

func (c *character) ExecuteModify(ctx *context.ModifierContext) {
	if ctx.AddLocalChargeModifiers() != nullChargeModifiers {
		localChargeModifiers := ctx.AddLocalChargeModifiers().Get(c.id)
		for _, localChargeModifier := range localChargeModifiers {
			c.localChargeModifiers.Append(localChargeModifier)
		}
	}

	if ctx.AddLocalHealModifiers() != nullHealModifiers {
		localHealModifiers := ctx.AddLocalHealModifiers().Get(c.id)
		for _, localHealModifier := range localHealModifiers {
			c.localHealModifiers.Append(localHealModifier)
		}
	}

	if ctx.AddLocalCostModifiers() != nullCostModifiers {
		localCostModifiers := ctx.AddLocalCostModifiers().Get(c.id)
		for _, localCostModifier := range localCostModifiers {
			c.localCostModifiers.Append(localCostModifier)
		}
	}

	if ctx.AddLocalDefenceModifiers() != nullDefenceModifiers {
		localDefenceModifiers := ctx.AddLocalDefenceModifiers().Get(c.id)
		for _, localDefenceModifier := range localDefenceModifiers {
			c.localDefenceModifiers.Append(localDefenceModifier)
		}
	}

	if ctx.AddLocalDirectAttackModifiers() != nullDirectAttackModifiers {
		localDirectAttackModifiers := ctx.AddLocalDirectAttackModifiers().Get(c.id)
		for _, localDirectAttackModifier := range localDirectAttackModifiers {
			c.localDirectAttackModifiers.Append(localDirectAttackModifier)
		}
	}

	if ctx.AddLocalFinalAttackModifiers() != nullFinalAttackModifiers {
		localFinalAttackModifiers := ctx.AddLocalFinalAttackModifiers().Get(c.id)
		for _, localFinalAttackModifier := range localFinalAttackModifiers {
			c.localFinalAttackModifiers.Append(localFinalAttackModifier)
		}
	}

	if ctx.RemoveLocalChargeModifiers() != nullChargeModifiers {
		localChargeModifiers := ctx.RemoveLocalChargeModifiers().Get(c.id)
		for _, localChargeModifier := range localChargeModifiers {
			c.localChargeModifiers.Remove(localChargeModifier.ID())
		}
	}

	if ctx.RemoveLocalHealModifiers() != nullHealModifiers {
		localHealModifiers := ctx.RemoveLocalHealModifiers().Get(c.id)
		for _, localHealModifier := range localHealModifiers {
			c.localHealModifiers.Remove(localHealModifier.ID())
		}
	}

	if ctx.RemoveLocalCostModifiers() != nullCostModifiers {
		localCostModifiers := ctx.RemoveLocalCostModifiers().Get(c.id)
		for _, localCostModifier := range localCostModifiers {
			c.localCostModifiers.Remove(localCostModifier.ID())
		}
	}

	if ctx.RemoveLocalDefenceModifiers() != nullDefenceModifiers {
		localDefenceModifiers := ctx.RemoveLocalDefenceModifiers().Get(c.id)
		for _, localDefenceModifier := range localDefenceModifiers {
			c.localDefenceModifiers.Remove(localDefenceModifier.ID())
		}
	}

	if ctx.RemoveLocalDirectAttackModifiers() != nullDirectAttackModifiers {
		localDirectAttackModifiers := ctx.RemoveLocalDirectAttackModifiers().Get(c.id)
		for _, localDirectAttackModifier := range localDirectAttackModifiers {
			c.localDirectAttackModifiers.Remove(localDirectAttackModifier.ID())
		}
	}

	if ctx.RemoveLocalFinalAttackModifiers() != nullFinalAttackModifiers {
		localFinalAttackModifiers := ctx.RemoveLocalFinalAttackModifiers().Get(c.id)
		for _, localFinalAttackModifier := range localFinalAttackModifiers {
			c.localFinalAttackModifiers.Remove(localFinalAttackModifier.ID())
		}
	}
}

func (c *character) ExecuteDefence(ctx *context.DamageContext) {
	c.localDefenceModifiers.Execute(ctx)
	executeAmount := ctx.Damage()[c.id]

	if executeAmount.Amount() >= c.currentHP {
		c.currentHP = 0
		c.status = enum.CharacterStatusDefeated
	} else {
		c.currentHP -= executeAmount.Amount()
	}
}

func (c *character) ExecuteAttack(skill, target uint, background []uint) (ctx *context.DamageContext) {
	s := c.skills.Get(skill)
	if attackSkill, ok := s.(AttackSkill); ok {
		return attackSkill.BaseDamage()
	} else {
		return context.NewEmptyDamageContext(skill, c.player, target, background)
	}
}

func (c *character) ExecuteDirectAttackModifiers(ctx *context.DamageContext) {
	c.localDirectAttackModifiers.Execute(ctx)
}

func (c *character) ExecuteFinalAttackModifiers(ctx *context.DamageContext) {
	c.localFinalAttackModifiers.Execute(ctx)
}

func (c character) ID() uint {
	return c.id
}

func (c character) Name() string {
	return c.name
}

func (c character) Description() string {
	return c.description
}

func (c character) Affiliation() enum.Affiliation {
	return c.affiliation
}

func (c character) Vision() enum.ElementType {
	return c.vision
}

func (c character) Weapon() enum.WeaponType {
	return c.weapon
}

func (c character) MaxHP() uint {
	return c.maxHP
}

func (c character) MaxMP() uint {
	return c.maxMP
}

func (c character) HP() uint {
	return c.currentHP
}

func (c character) MP() uint {
	return c.currentMP
}

func (c character) Status() enum.CharacterStatus {
	return c.status
}

func NewCharacter() Character {
	return &character{}
}
