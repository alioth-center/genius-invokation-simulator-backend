package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type Map struct {
	sets kv.Map[enum.TriggerType, *set]
}

// Call 调用某种类型的Event
func (m *Map) Call(triggerType enum.TriggerType, ctx *context.CallbackContext) {
	if m.sets.Exists(triggerType) {
		set := m.sets.Get(triggerType)
		set.call(ctx)
	}
}

// Preview 预览某种类型的Event
func (m Map) Preview(triggerType enum.TriggerType, ctx *context.CallbackContext) {
	if m.sets.Exists(triggerType) {
		set := m.sets.Get(triggerType)
		set.preview(ctx)
	}
}

// AddEvent 添加一个Event
func (m *Map) AddEvent(event Event) {
	if m.sets.Exists(event.TriggerAt()) {
		m.sets.Get(event.TriggerAt()).add(event)
	} else {
		set := newEventSet()
		set.add(event)
		m.sets.Set(event.TriggerAt(), set)
	}
}

// RemoveEvent 移除一个Event
func (m *Map) RemoveEvent(event Event) {
	if m.sets.Exists(event.TriggerAt()) {
		m.sets.Get(event.TriggerAt()).remove(event.ID())
	}
}

// AddEvents 添加多个Event，只有类型为filter的Event会生效
func (m *Map) AddEvents(filter enum.TriggerType, events []Event) {
	set := newEventSet()
	for _, event := range events {
		if event.TriggerAt() == filter {
			set.add(event)
		}
	}

	if m.sets.Exists(filter) {
		m.sets.Get(filter).append(set)
	} else {
		m.sets.Set(filter, set)
	}
}

// RemoveEvents 移除类型为filter的Event
func (m *Map) RemoveEvents(filter enum.TriggerType) {
	if m.sets.Exists(filter) {
		m.sets.Set(filter, newEventSet())
	}
}

func NewEventMap() *Map {
	return &Map{sets: kv.NewCommonMap[enum.TriggerType, *set]()}
}
