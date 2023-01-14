package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Event interface {
	BaseEntity
	TriggerAt() enum.TriggerType
	CanTriggered(context.CallbackContext) bool
	NeedClear() bool
	Callback(*context.CallbackContext)
}
