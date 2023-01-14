package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Map struct {
	sets map[enum.TriggerType]*set
}

// Call 调用某种类型的Event
func (m *Map) Call(triggerType enum.TriggerType, ctx *context.CallbackContext) {
	if set, exist := m.sets[triggerType]; exist {
		set.call(ctx)
	}
}

// Preview 预览某种类型的Event
func (m Map) Preview(triggerType enum.TriggerType, ctx *context.CallbackContext) {
	if set, exist := m.sets[triggerType]; exist {
		set.preview(ctx)
	}
}

// AddEvent 添加一个Event
func (m *Map) AddEvent(event Event) {
	if set, exist := m.sets[event.TriggerAt()]; !exist || set == nil {
		set = newEventSet()
		set.add(event)
		m.sets[event.TriggerAt()] = set
	} else {
		m.sets[event.TriggerAt()].add(event)
	}
}

// RemoveEvent 移除一个Event
func (m *Map) RemoveEvent(event Event) {
	if set, exist := m.sets[event.TriggerAt()]; exist && set != nil {
		set.remove(event.ID())
	}
}

// AddEvents 添加多个Event，只有类型为filter的Event会生效
func (m *Map) AddEvents(filter enum.TriggerType, events []Event) {
	if set, exist := m.sets[filter]; !exist || set == nil {
		set = newEventSet()
		for _, event := range events {
			if event.TriggerAt() == filter {
				set.add(event)
			}
		}
		m.sets[filter] = set
	} else {
		for _, event := range events {
			if event.TriggerAt() == filter {
				set.add(event)
			}
		}
	}
}

// RemoveEvents 移除类型为filter的Event
func (m *Map) RemoveEvents(filter enum.TriggerType) {
	delete(m.sets, filter)
}

func (m Map) Expose(trigger enum.TriggerType) (events []uint) {
	if set, exist := m.sets[trigger]; !exist {
		return []uint{}
	} else {
		events = []uint{}
		for id := range set.events {
			events = append(events, id)
		}
		return events
	}
}

func NewEventMap() *Map {
	return &Map{sets: map[enum.TriggerType]*set{}}
}

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
