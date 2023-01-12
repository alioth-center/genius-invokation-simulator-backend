package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type set struct {
	events map[uint]Event
}

// call 调用EventSet中的所有Event，并在调用完成后清理需要清理的Event
func (s *set) call(ctx *context.CallbackContext) {
	for _, e := range s.events {
		if e.CanTriggered(*ctx) {
			e.Callback(ctx)
			delete(s.events, e.ID())
		}
	}
}

// preview 调用EventSet中的所有Event，但调用完成后不清理调用过的Event
func (s set) preview(ctx *context.CallbackContext) {
	for _, e := range s.events {
		if e.CanTriggered(*ctx) {
			e.Callback(ctx)
		}
	}
}

// append 合并两个EventSet，若新Set中有同id的Event，则会覆盖现有Event
func (s *set) append(another *set) {
	for id, event := range another.events {
		s.events[id] = event
	}
}

// add 向EventSet中加入一个Event
func (s *set) add(event Event) {
	s.events[event.ID()] = event
}

// remove 从EventSet中移除一个指定id的Event
func (s *set) remove(id uint) {
	delete(s.events, id)
}

// newEventSet 创建一个空EventSet
func newEventSet() *set {
	return &set{events: map[uint]Event{}}
}
