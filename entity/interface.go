package entity

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

// Player 供外部包调用的玩家接口
type Player interface {
	GetUID() (uid uint)
	GetCost() (cost map[enum.ElementType]uint)
	GetCards() (cards []uint)
	GetSummons() (summons []uint)
	GetSupports() (supports []uint)
	CardDeckRemain() (remain uint)
	GetActiveCharacter() (character uint)
	GetBackgroundCharacters() (characters []uint)
	GetCharacter(character uint) (entity Character)
	GetStatus() (status enum.PlayerStatus)
	GetGlobalModifiers(modifierType enum.ModifierType) (modifiers []uint)
	GetCooperativeSkills(trigger enum.TriggerType) (skills []uint)
	GetEvents(trigger enum.TriggerType) (events []uint)
}

// Character 供外部包调用的角色接口
type Character interface {
	GetID() (id uint)
	GetOwner() (owner uint)
	GetAffiliation() (affiliation enum.Affiliation)
	GetVision() (element enum.ElementType)
	GetWeaponType() (weaponType enum.WeaponType)
	GetSkills() (skills []uint)
	GetHP() (hp uint)
	GetMaxHP() (maxHP uint)
	GetMP() (mp uint)
	GetMaxMP() (maxMP uint)
	GetEquipment(equipmentType enum.EquipmentType) (equipped bool, equipment uint)
	GetSatiety() (satiety bool)
	GetAttachedElements() (elements []enum.ElementType)
	GetStatus() (status enum.CharacterStatus)
	GetLocalModifiers(modifierType enum.ModifierType) (modifiers []uint)
}
