package event

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type Event interface {
	Info() Info
}

type Info struct {
	TriggerAt enum.TriggerType
}
