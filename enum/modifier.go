package enum

type ModifierType byte

const (
	ModifierTypeNone ModifierType = iota
	ModifierTypeAttack
	ModifierTypeCharacter
	ModifierTypeCharge
	ModifierTypeCost
	ModifierTypeDefence
	ModifierTypeHeal
)
