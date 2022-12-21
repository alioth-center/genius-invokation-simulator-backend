package enum

// CharacterStatus 角色的状态
type CharacterStatus byte

const (
	CharacterStatusReady      CharacterStatus = iota // CharacterStatusReady 角色已就绪，无状态
	CharacterStatusActive                            // CharacterStatusActive 角色已激活，前台角色
	CharacterStatusBackground                        // CharacterStatusBackground 角色已激活，后台角色
	CharacterStatusDisabled                          // CharacterStatusDisabled 前台角色，但无法进行操作
	CharacterStatusDefeated                          // CharacterStatusDefeated 角色已被击败，不可操作
)
