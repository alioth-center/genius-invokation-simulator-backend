package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"testing"
)

func TestCostEquals(t *testing.T) {
	tests := []struct {
		name   string
		origin map[enum.ElementType]uint
		cost   map[enum.ElementType]uint
		want   bool
	}{
		{
			name:   "ElementSetEquals-Common-1",
			origin: map[enum.ElementType]uint{enum.ElementPyro: 3},
			cost:   map[enum.ElementType]uint{enum.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Common-2",
			origin: map[enum.ElementType]uint{enum.ElementPyro: 2, enum.ElementDendro: 1},
			cost:   map[enum.ElementType]uint{enum.ElementPyro: 2, enum.ElementDendro: 1},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-1",
			origin: map[enum.ElementType]uint{enum.ElementCurrency: 3},
			cost:   map[enum.ElementType]uint{enum.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-2",
			origin: map[enum.ElementType]uint{enum.ElementCurrency: 2, enum.ElementPyro: 1},
			cost:   map[enum.ElementType]uint{enum.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-3",
			origin: map[enum.ElementType]uint{enum.ElementCurrency: 2, enum.ElementCryo: 1},
			cost:   map[enum.ElementType]uint{enum.ElementPyro: 3},
			want:   false,
		},
		{
			name:   "ElementSetEquals-Mixed-1",
			origin: map[enum.ElementType]uint{enum.ElementCurrency: 2, enum.ElementCryo: 1},
			cost:   map[enum.ElementType]uint{enum.ElementNone: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-2",
			origin: map[enum.ElementType]uint{enum.ElementCryo: 1, enum.ElementCurrency: 1},
			cost:   map[enum.ElementType]uint{enum.ElementCurrency: 2},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-3",
			origin: map[enum.ElementType]uint{enum.ElementCryo: 1, enum.ElementDendro: 1, enum.ElementAnemo: 1},
			cost:   map[enum.ElementType]uint{enum.ElementCurrency: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-4",
			origin: map[enum.ElementType]uint{enum.ElementCryo: 2, enum.ElementDendro: 2, enum.ElementCurrency: 1},
			cost:   map[enum.ElementType]uint{enum.ElementCurrency: 5},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originCost := newCostFromMap(tt.origin)
			otherCost := newCostFromMap(tt.cost)
			if result := originCost.Equals(*otherCost); result != tt.want {
				t.Errorf("incorrect result, want %v, got %v", tt.want, result)
			}
		})
	}
}

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
	originCost := newCostFromMap(tt.origin)
	otherCost := newCostFromMap(tt.cost)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if got := originCost.Equals(*otherCost); got != tt.want {
			b.Errorf("Error: %v, want %v", got, tt.want)
		}
	}
}
