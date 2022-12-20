package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Event interface {
	Info() Info
	Triggered(context.CallbackContext) bool
	Callback() func(*context.CallbackContext)
}

type Info struct {
	triggerAt   enum.TriggerType
	name        string
	description string
}

func (i Info) TriggerAt() enum.TriggerType {
	return i.triggerAt
}

func (i Info) Name() string {
	return i.name
}

func (i Info) Description() string {
	return i.description
}
