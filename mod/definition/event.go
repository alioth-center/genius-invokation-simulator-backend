package definition

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Event interface {
	TriggerAt() enum.TriggerType
	TriggeredNow(ctx context.CallbackContext) bool
	ClearNow() bool
	CallBack(callbackContext *context.CallbackContext)
}
