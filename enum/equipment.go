package enum

// EquipmentType 装备类型
type EquipmentType byte

const (
	EquipmentNone     EquipmentType = iota // EquipmentNone 无装备
	EquipmentWeapon                        // EquipmentWeapon 武器
	EquipmentArtifact                      // EquipmentArtifact 圣遗物
	EquipmentTalent                        // EquipmentTalent 天赋
)
