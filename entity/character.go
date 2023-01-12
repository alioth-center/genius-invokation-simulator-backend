package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type character struct {
	id          uint             // id 角色的ID，由框架确定
	player      uint             // player 所属玩家的ID，由框架确定
	affiliation enum.Affiliation // affiliation 角色的势力归属
	vision      enum.ElementType // vision 角色的元素类型
	weapon      enum.WeaponType  // weapon 角色的武器类型
	skills      map[uint]Skill   // skills 角色的技能

	maxHP      uint                        // maxHP 角色的最大生命值
	currentHP  uint                        // currentHP 角色的当前生命值
	maxMP      uint                        // maxMP 角色的最大能量值
	currentMP  uint                        // currentMP 角色的当前能量值
	status     enum.CharacterStatus        // status 角色的状态
	elements   []enum.ElementType          // elements 角色目前附着的元素
	satiety    bool                        // satiety 角色的饱腹状态
	equipments map[enum.EquipmentType]uint // equipments 角色穿着的装备

	localDirectAttackModifiers AttackModifiers  // localDirectAttackModifiers 本地直接攻击修正
	localFinalAttackModifiers  AttackModifiers  // localFinalAttackModifiers 本地最终攻击修正
	localDefenceModifiers      DefenceModifiers // localDefenceModifiers 本地防御修正
	localChargeModifiers       ChargeModifiers  // localChargeModifiers 本地充能修正
	localHealModifiers         HealModifiers    // localHealModifiers 本地治疗修正
	localCostModifiers         CostModifiers    // localCostModifiers 本地费用修正

	ruleSet RuleSet // ruleSet 用于结算的规则集合
}

func (c character) GetID() (id uint) {
	return c.id
}

func (c character) GetOwner() (owner uint) {
	return c.player
}

func (c character) GetAffiliation() (affiliation enum.Affiliation) {
	return c.affiliation
}

func (c character) GetVision() (element enum.ElementType) {
	return c.vision
}

func (c character) GetWeaponType() (weaponType enum.WeaponType) {
	return c.weapon
}

func (c character) GetSkills() (skills []uint) {
	skills = make([]uint, 0)
	for id := range c.skills {
		skills = append(skills, id)
	}
	return skills
}

func (c character) GetHP() (hp uint) {
	return c.currentHP
}

func (c character) GetMaxHP() (maxHP uint) {
	return c.maxHP
}

func (c character) GetMP() (mp uint) {
	return c.currentMP
}

func (c character) GetMaxMP() (maxMP uint) {
	return c.maxMP
}

func (c character) GetEquipment(equipmentType enum.EquipmentType) (equipped bool, equipment uint) {
	equipmentID, exist := c.equipments[equipmentType]
	return exist, equipmentID
}

func (c character) GetSatiety() (satiety bool) {
	return c.satiety
}

func (c character) GetAttachedElements() (elements []enum.ElementType) {
	return c.elements
}

func (c character) GetStatus() (status enum.CharacterStatus) {
	return c.status
}

func (c character) GetLocalModifiers(modifierType enum.ModifierType) (modifiers []uint) {
	switch modifierType {
	case enum.ModifierTypeNone:
		return []uint{}
	case enum.ModifierTypeAttack:
		modifiers = []uint{}
		modifiers = append(modifiers, c.localDirectAttackModifiers.Expose()...)
		modifiers = append(modifiers, c.localFinalAttackModifiers.Expose()...)
		return modifiers
	case enum.ModifierTypeCharacter:
		return []uint{}
	case enum.ModifierTypeCharge:
		return c.localChargeModifiers.Expose()
	case enum.ModifierTypeCost:
		return c.localCostModifiers.Expose()
	case enum.ModifierTypeDefence:
		return c.localDefenceModifiers.Expose()
	case enum.ModifierTypeHeal:
		return c.localHealModifiers.Expose()
	default:
		return []uint{}
	}
}
