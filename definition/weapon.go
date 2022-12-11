package definition

// Weapon 武器类型
type Weapon uint

const (
	WeaponSword    Weapon = iota // WeaponSword 单手剑
	WeaponClaymore               // WeaponClaymore 双手剑
	WeaponBow                    // WeaponBow 弓
	WeaponCatalyst               // WeaponCatalyst 法器
	WeaponPolearm                // WeaponPolearm 长柄武器
	WeaponOthers                 // WeaponOthers 其他武器
)
