package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"testing"
)

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
