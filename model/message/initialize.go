package message

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type MatchingMessage struct {
	UID        uint
	Characters []uint
	CardDeck   []uint
}

type GameOptions struct {
	ReRollTime    uint
	ElementAmount uint
	GetCards      uint
	StaticElement map[enum.ElementType]uint
	RuleSet       uint
}

type InitializeMessage struct {
	Players []MatchingMessage
	Options GameOptions
}
