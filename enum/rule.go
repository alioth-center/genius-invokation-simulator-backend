package enum

type RuleType byte

const (
	RuleTypeNone RuleType = iota
	RuleTypeReactionCalculator
	RuleTypeVictorCalculator
)
