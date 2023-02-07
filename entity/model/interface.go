package model

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

// Player 供外部包调用的玩家接口
type Player interface {
	GetUID() (uid uint64)
	GetCost() (cost map[enum.ElementType]uint)
	GetCards() (cards []uint64)
	GetSummons() (summons []uint64)
	GetSupports() (supports []uint64)
	CardDeckRemain() (remain uint)
	GetActiveCharacter() (character uint64)
	GetBackgroundCharacters() (characters []uint64)
	GetCharacter(character uint64) (has bool, entity Character)
	GetStatus() (status enum.PlayerStatus)
	GetGlobalModifiers(modifierType enum.ModifierType) (modifiers []uint64)
	GetCooperativeSkills(trigger enum.TriggerType) (skills []uint64)
	GetEvents(trigger enum.TriggerType) (events []uint64)
}

// Character 供外部包调用的角色接口
type Character interface {
	GetID() (id uint64)
	GetOwner() (owner uint64)
	GetAffiliation() (affiliation enum.Affiliation)
	GetVision() (element enum.ElementType)
	GetWeaponType() (weaponType enum.WeaponType)
	GetSkills() (skills []uint64)
	GetHP() (hp uint)
	GetMaxHP() (maxHP uint)
	GetMP() (mp uint)
	GetMaxMP() (maxMP uint)
	GetEquipment(equipmentType enum.EquipmentType) (equipped bool, equipment uint64)
	GetSatiety() (satiety bool)
	GetAttachedElements() (elements []enum.ElementType)
	GetStatus() (status enum.CharacterStatus)
	GetLocalModifiers(modifierType enum.ModifierType) (modifiers []uint64)
}
