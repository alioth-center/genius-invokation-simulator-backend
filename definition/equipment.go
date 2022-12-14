/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "equipment.go" LastUpdatedAt 2022/12/14 09:15:14
 */

package definition

type EquipmentType byte

const (
	EquipmentNone     EquipmentType = iota // EquipmentNone 无装备
	EquipmentWeapon                        // EquipmentWeapon 武器
	EquipmentArtifact                      // EquipmentArtifact 圣遗物
	EquipmentTalent                        // EquipmentTalent 天赋
)
