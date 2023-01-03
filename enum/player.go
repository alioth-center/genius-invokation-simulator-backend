package enum

type PlayerStatus byte

const (
	PlayerStatusReady PlayerStatus = iota
	PlayerStatusWaiting
	PlayerStatusActing
	PlayerStatusDefeated
	PlayerStatusViewing
)
