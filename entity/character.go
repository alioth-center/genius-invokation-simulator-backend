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

type CharacterInfo interface {
	ID() uint
	Affiliation() enum.Affiliation
	Vision() enum.ElementType
	Weapon() enum.WeaponType
	MaxHP() uint
	MaxMP() uint
	Skills() map[uint]Skill
}

type Character interface {
	// ID 角色的ID
	ID() uint

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

	// HasSkill 角色是否持有id为skill的技能
	HasSkill(skill uint) bool

	// Status 角色的状态
	Status() enum.CharacterStatus

	// SwitchUp 切换到前台
	SwitchUp()

	// SwitchDown 切换到后台
	SwitchDown()

	// ExecuteCharge 根据ChargeContext给角色增加或减少MP
	ExecuteCharge(ctx *context.ChargeContext)

	// ExecuteHeal 根据HealContext给角色进行治疗
	ExecuteHeal(ctx *context.HealContext)

	// PreviewCostModify 预览CostModifiers的效果
	PreviewCostModify(ctx *context.CostContext)

	// ExecuteCostModify 使用角色的CostModifiers对CostContext进行修正
	ExecuteCostModify(ctx *context.CostContext)

	// ExecuteModify 根据ModifierContext对角色的Modifiers进行修改
	ExecuteModify(ctx *context.ModifierContext)

	// ExecuteDefence 根据DamageContext对角色进行伤害结算，不包括效果结算
	ExecuteDefence(ctx *context.DamageContext)

	// ExecuteEatFood 根据食物卡的效果执行食物结算
	ExecuteEatFood(ctx *context.ModifierContext)

	// ExecuteAttack 使用skill进行对target角色和background后台角色进行攻击
	ExecuteAttack(skill, target uint, background []uint) (ctx *context.DamageContext)

	// ExecuteDirectAttackModifiers 使用角色的DirectAttackModifiers对DamageContext进行伤害修正
	ExecuteDirectAttackModifiers(ctx *context.DamageContext)

	// ExecuteFinalAttackModifiers 使用角色的FinalAttackModifiers对DamageContext进行伤害修正
	ExecuteFinalAttackModifiers(ctx *context.DamageContext)

	// ExecuteElementAttachment 判断角色能否附着attachElement元素并尝试进行附着，此时不触发元素反应
	ExecuteElementAttachment(attachElement enum.ElementType)

	// ExecuteElementReaction 尝试使用角色身上附着的元素进行反应，返回能否反应和反应类型
	ExecuteElementReaction() (reaction enum.Reaction)
}

type character struct {
	id          uint                // id 角色的ID，由框架确定
	player      uint                // player 所属玩家的ID，由框架确定
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

	ruleSet RuleSet // ruleSet 用于结算的规则集合
}

func (c *character) SwitchUp() {
	c.status = enum.CharacterStatusActive
}

func (c *character) SwitchDown() {
	c.status = enum.CharacterStatusBackground
}

func (c character) HasSkill(skill uint) bool {
	return c.skills.Exists(skill)
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
		if c.currentMP < uint(executeAmount) {
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

func (c *character) PreviewCostModify(ctx *context.CostContext) {
	c.localCostModifiers.Preview(ctx)
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
		return attackSkill.BaseDamage(target, c.player, background)
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

func (c *character) ExecuteEatFood(ctx *context.ModifierContext) {
	c.ExecuteModify(ctx)
	c.satiety = true
}

func (c *character) ExecuteElementAttachment(attachElement enum.ElementType) {
	c.elements = c.ruleSet.ReactionCalculator().Attach(c.elements, attachElement)
}

func (c *character) ExecuteElementReaction() (reaction enum.Reaction) {
	reaction, c.elements = c.ruleSet.ReactionCalculator().ReactionCalculate(c.elements)
	return reaction
}

func (c character) ID() uint {
	return c.id
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

func NewCharacter(owner uint, info CharacterInfo, ruleSet RuleSet) Character {
	character := &character{
		id:                         info.ID(),
		player:                     owner,
		affiliation:                info.Affiliation(),
		vision:                     info.Vision(),
		weapon:                     info.Weapon(),
		skills:                     kv.NewSimpleMap[Skill](),
		maxHP:                      info.MaxHP(),
		currentHP:                  info.MaxHP(),
		maxMP:                      info.MaxMP(),
		currentMP:                  0,
		status:                     enum.CharacterStatusReady,
		elements:                   []enum.ElementType{},
		satiety:                    false,
		equipments:                 kv.NewSimpleMap[interface{}](),
		localDirectAttackModifiers: modifier.NewChain[context.DamageContext](),
		localFinalAttackModifiers:  modifier.NewChain[context.DamageContext](),
		localDefenceModifiers:      modifier.NewChain[context.DamageContext](),
		localChargeModifiers:       modifier.NewChain[context.ChargeContext](),
		localHealModifiers:         modifier.NewChain[context.HealContext](),
		localCostModifiers:         modifier.NewChain[context.CostContext](),
		ruleSet:                    ruleSet,
	}

	for id, skill := range info.Skills() {
		character.skills.Set(id, skill)
	}

	return character
}
