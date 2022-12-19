package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type Character struct {
	id          uint                 // id 角色的UID，由框架确定
	name        string               // name 角色的名称
	description string               // description 角色的描述
	affiliation enum.Affiliation     // affiliation 角色的势力归属
	vision      enum.ElementType     // vision 角色的元素类型
	weapon      enum.WeaponType      // weapon 角色的武器类型
	skills      map[uint]interface{} // skills 角色的技能

	maxHP      uint                               // maxHP 角色的最大生命值
	currentHP  uint                               // currentHP 角色的当前生命值
	maxMP      uint                               // maxMP 角色的最大能量值
	currentMP  uint                               // currentMP 角色的当前能量值
	status     enum.CharacterStatus               // status 角色的状态
	elements   []byte                             // elements 角色目前附着的元素
	satiety    bool                               // satiety 角色的饱腹状态
	equipments map[enum.EquipmentType]interface{} // equipments 角色穿着的装备

	localAttackModifiers  modifier.Chain[context.DamageContext] // localAttackModifiers 本地攻击修正
	localDefenceModifiers modifier.Chain[context.DamageContext] // localDefenceModifiers 本地防御修正
	localChargeModifiers  modifier.Chain[context.ChargeContext] // localChargeModifiers 本地充能修正
	localHealModifiers    modifier.Chain[context.HealContext]   // localHealModifiers 本地治疗修正
	localCostModifiers    modifier.Chain[context.CostContext]   // localCostModifiers 本地费用修正
	callbackEvents        interface{}                           // callbackEvents 回调事件
}

func (c Character) ID() uint {
	return c.id
}

func (c Character) Name() string {
	return c.name
}

func (c Character) Description() string {
	return c.description
}

func (c Character) Affiliation() enum.Affiliation {
	return c.affiliation
}

func (c Character) Vision() enum.ElementType {
	return c.vision
}

func (c Character) Weapon() enum.WeaponType {
	return c.weapon
}

func (c Character) MaxHP() uint {
	return c.maxHP
}

func (c Character) MaxMP() uint {
	return c.maxMP
}

func (c Character) HP() uint {
	return c.currentHP
}

func (c Character) MP() uint {
	return c.currentMP
}

func (c Character) Status() enum.CharacterStatus {
	return c.status
}
