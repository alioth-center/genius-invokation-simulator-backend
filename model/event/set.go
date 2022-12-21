package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type set struct {
	events kv.Map[uint, Event]
}

// call 调用EventSet中的所有Event，并在调用完成后清理调用过的Event
func (s *set) call(ctx *context.CallbackContext) {
	s.events.Range(func(id uint, event Event) bool {
		if event.CanTriggered(*ctx) {
			event.Callback(ctx)
			s.events.Remove(id)
		}
		return true
	})
}

// preview 调用EventSet中的所有Event，但调用完成后不清理调用过的Event
func (s set) preview(ctx *context.CallbackContext) {
	s.events.Range(func(id uint, event Event) bool {
		if event.CanTriggered(*ctx) {
			event.Callback(ctx)
		}
		return true
	})
}

// append 合并两个EventSet，若新Set中有同id的Event，则会覆盖现有Event
func (s *set) append(another *set) {
	another.events.Range(func(k uint, v Event) bool {
		s.events.Set(k, v)
		return true
	})
}

// add 向EventSet中加入一个Event
func (s *set) add(event Event) {
	s.events.Set(event.ID(), event)
}

// remove 从EventSet中移除一个指定id的Event
func (s *set) remove(id uint) {
	s.events.Remove(id)
}

// newEventSet 创建一个空EventSet
func newEventSet() *set {
	return &set{events: kv.NewSimpleMap[Event]()}
}
