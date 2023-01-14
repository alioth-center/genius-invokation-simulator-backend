package event

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
