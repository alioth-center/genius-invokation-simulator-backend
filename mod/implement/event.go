package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type EventImpl struct {
	EntityImpl
	triggerAt  enum.TriggerType
	triggerNow func(ctx context.CallbackContext) bool
	clearNow   func() bool
	callback   func(ctx *context.CallbackContext)
}

func (impl *EventImpl) TriggerAt() enum.TriggerType {
	return impl.triggerAt
}

func (impl *EventImpl) TriggeredNow(ctx context.CallbackContext) bool {
	return impl.triggerNow(ctx)
}

func (impl *EventImpl) ClearNow() bool {
	return impl.clearNow()
}

func (impl *EventImpl) CallBack(callbackContext *context.CallbackContext) {
	impl.callback(callbackContext)
}

type EventOptions func(option *EventImpl)

func WithEventID(id uint16) EventOptions {
	return func(option *EventImpl) {
		option.InjectTypeID(uint64(id))
	}
}

func WithEventTriggerAt(triggerAt enum.TriggerType) EventOptions {
	return func(option *EventImpl) {
		option.triggerAt = triggerAt
	}
}

func WithEventTriggerNow(triggerNow func(ctx context.CallbackContext) bool) EventOptions {
	return func(option *EventImpl) {
		option.triggerNow = triggerNow
	}
}

func WithEventClearNow(clearNow func() bool) EventOptions {
	return func(option *EventImpl) {
		option.clearNow = clearNow
	}
}

func WithEventCallback(callback func(ctx *context.CallbackContext)) EventOptions {
	return func(option *EventImpl) {
		option.callback = callback
	}
}

func NewEventWithOpts(options ...EventOptions) definition.Event {
	impl := &EventImpl{}
	for _, option := range options {
		option(impl)
	}

	return impl
}
