/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "character-status.go" LastUpdatedAt 2022/12/12 10:19:12
 */

package definition

// CharacterStatus 角色的状态
type CharacterStatus byte

const (
	CharacterStatusReady        CharacterStatus = iota // CharacterStatusReady 角色已就绪，无状态
	CharacterStatusActive                              // CharacterStatusActive 角色已激活，前台角色
	CharacterStatusBackground                          // CharacterStatusBackground 角色已激活，后台角色
	CharacterStatusDisabled                            // CharacterStatusDisabled 前台角色，但无法进行操作
	CharacterStatusUnselectable                        // CharacterStatusUnselectable 角色不可操作
)
