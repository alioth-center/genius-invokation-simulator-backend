package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
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
	originCost := NewCostFromMap(tt.origin)
	otherCost := NewCostFromMap(tt.cost)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if got := originCost.Equals(*otherCost); got != tt.want {
			b.Errorf("Error: %v, want %v", got, tt.want)
		}
	}
}
