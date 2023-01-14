package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"testing"
)

func BenchmarkTestCostEquals(b *testing.B) {
	tt := struct {
		name   string
		origin map[enum.ElementType]uint
		cost   map[enum.ElementType]uint
		want   bool
	}{
		name:   "ElementSetEquals-Mixed-4",
		origin: map[enum.ElementType]uint{enum.ElementCryo: 2, enum.ElementDendro: 2, enum.ElementCurrency: 1},
		cost:   map[enum.ElementType]uint{enum.ElementCurrency: 3, enum.ElementNone: 2},
		want:   true,
	}
	originCost := model.NewCostFromMap(tt.origin)
	otherCost := model.NewCostFromMap(tt.cost)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if got := originCost.Equals(*otherCost); got != tt.want {
			b.Errorf("Error: %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkTestPlayerChainNext(b *testing.B) {
	pc := newPlayerChain()
	for i := uint(0); i < 100; i++ {
		pc.add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pc.next()
	}
}

func BenchmarkTestPlayerChainNextWithComplete(b *testing.B) {
	pc := newPlayerChain()
	for i := uint(0); i < 11451419; i++ {
		pc.add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exist, _ := pc.next()
		pc.complete(uint(i))
		if !exist {
			break
		}
	}
}

func BenchmarkTestEventMapPreview(b *testing.B) {
	m := model.NewEventMap()
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
	m := model.NewEventMap()
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
	m := model.NewEventMap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.AddEvent(newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) }))
	}
}

func BenchmarkTestEventMapRemove(b *testing.B) {
	m := model.NewEventMap()
	event := newTestEvent(uint(1), enum.AfterSwitch, true, func(ctx *context.CallbackContext) { ctx.ChangeOperated(false) })
	m.AddEvent(event)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.RemoveEvent(event)
	}
}
