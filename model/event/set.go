package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type Set struct {
	events kv.Map[uint, Event]
}

// Call 调用EventSet中的所有Event，并在调用完成后清理调用过的Event
func (s *Set) Call(ctx *context.CallbackContext) {
	s.events.Range(func(id uint, event Event) bool {
		if event.Triggered(*ctx) {
			event.Callback()(ctx)
			defer s.events.Remove(id)
		}
		return true
	})
}

// Append 合并两个EventSet，若新Set中有同id的Event，则会覆盖现有Event
func (s *Set) Append(another Set) {
	another.events.Range(func(k uint, v Event) bool {
		s.events.Set(k, v)
		return true
	})
}
