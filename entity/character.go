package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type Character interface {
}

type character struct {
	id          uint                // id 角色的ID，由框架确定
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

func (c *character) Modify(modifyFunc func(ctx *context.ModifierContext)) {
	self := &context.ModifierContext{}
	modifyFunc(self)
}

func NewCharacter() Character {
	return &character{}
}
