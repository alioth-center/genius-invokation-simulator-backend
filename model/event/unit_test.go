package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"testing"
)

type testEvent struct {
	id      uint
	trigger enum.TriggerType
	enable  bool
	handler func(*context.CallbackContext)
}

func (t testEvent) ID() uint { return t.id }

func (t testEvent) TriggerAt() enum.TriggerType { return t.trigger }

func (t testEvent) CanTriggered(callbackContext context.CallbackContext) bool { return t.enable }

func (t testEvent) Callback(ctx *context.CallbackContext) { t.handler(ctx) }

func newTestEvent(id uint, trigger enum.TriggerType, enable bool, handler func(callbackContext *context.CallbackContext)) Event {
	return &testEvent{id: id, trigger: trigger, enable: enable, handler: handler}
}

func TestEventMapExecute(t *testing.T) {
	switchCharacterEnabled := newTestEvent(0, enum.AfterAttack, true, func(ctx *context.CallbackContext) { ctx.SwitchCharacter(114514) })
	changeOperatedEnabled := newTestEvent(1, enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) })
	switchCharacterDisabled := newTestEvent(2, enum.AfterAttack, false, func(ctx *context.CallbackContext) { ctx.SwitchCharacter(114514) })
	changeOperatedDisabled := newTestEvent(3, enum.AfterSwitch, false, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) })

	testCases := []struct {
		name      string
		triggers  []Event
		call      enum.TriggerType
		judgeFunc func(*context.CallbackContext) bool
		want      bool
	}{
		{
			name:     "TestEventMapExecute-1",
			triggers: []Event{switchCharacterEnabled, changeOperatedEnabled},
			call:     enum.AfterAttack,
			judgeFunc: func(ctx *context.CallbackContext) bool {
				status, target := ctx.SwitchCharacterResult()
				return status && target == 114514
			},
			want: true,
		},
		{
			name:     "TestEventMapExecute-2",
			triggers: []Event{switchCharacterDisabled, changeOperatedDisabled},
			call:     enum.AfterAttack,
			judgeFunc: func(ctx *context.CallbackContext) bool {
				status, target := ctx.SwitchCharacterResult()
				return status && target == 114514
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			m := NewEventMap()
			for _, trigger := range tt.triggers {
				m.AddEvent(trigger)
			}

			ctx := context.NewCallbackContext()
			m.Call(tt.call, ctx)
			if got := tt.judgeFunc(ctx); got != tt.want {
				t.Errorf("incorrect execute result: got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTestEventMapPreview(b *testing.B) {
	m := NewEventMap()
	for i := 0; i < 128; i++ {
		m.AddEvent(newTestEvent(uint(i), enum.AfterAttack, true, func(ctx *context.CallbackContext) { ctx.SwitchCharacter(114514) }))
	}
	m.AddEvent(newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) }))
	ctx := context.NewCallbackContext()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Preview(enum.AfterSwitch, ctx)
	}
}

func BenchmarkTestEventMapExecute(b *testing.B) {
	m := NewEventMap()
	for i := 0; i < 128; i++ {
		m.AddEvent(newTestEvent(uint(i), enum.AfterAttack, true, func(ctx *context.CallbackContext) { ctx.SwitchCharacter(114514) }))
	}
	m.AddEvent(newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) }))
	ctx := context.NewCallbackContext()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Call(enum.AfterSwitch, ctx)
	}
}

func BenchmarkTestEventMapAdd(b *testing.B) {
	m := NewEventMap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.AddEvent(newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) }))
	}
}

func BenchmarkTestEventMapRemove(b *testing.B) {
	m := NewEventMap()
	event := newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) })
	m.AddEvent(event)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.RemoveEvent(event)
	}
}
