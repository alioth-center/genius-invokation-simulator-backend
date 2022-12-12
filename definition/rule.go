/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "rule.go" LastUpdatedAt 2022/12/12 13:59:12
 */

package definition

type RuleType byte

const (
	RuleInitializeGame RuleType = iota
	RuleInGameModifier
)
