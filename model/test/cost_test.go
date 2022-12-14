/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "cost_test.go" LastUpdatedAt 2022/12/14 16:07:14
 */

package test

import (
	d "github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"
	"testing"
)

func TestElementSetEquals(t *testing.T) {
	tests := []struct {
		name   string
		origin d.ElementSet
		cost   d.ElementSet
		want   bool
	}{
		{
			name:   "ElementSetEquals-Common-1",
			origin: map[d.Element]uint{d.ElementPyro: 3},
			cost:   map[d.Element]uint{d.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Common-2",
			origin: map[d.Element]uint{d.ElementPyro: 2, d.ElementDendro: 1},
			cost:   map[d.Element]uint{d.ElementPyro: 2, d.ElementDendro: 1},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-1",
			origin: map[d.Element]uint{d.ElementCurrency: 3},
			cost:   map[d.Element]uint{d.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-2",
			origin: map[d.Element]uint{d.ElementCurrency: 2, d.ElementPyro: 1},
			cost:   map[d.Element]uint{d.ElementPyro: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Currency-3",
			origin: map[d.Element]uint{d.ElementCurrency: 2, d.ElementCryo: 1},
			cost:   map[d.Element]uint{d.ElementPyro: 3},
			want:   false,
		},
		{
			name:   "ElementSetEquals-Mixed-1",
			origin: map[d.Element]uint{d.ElementCurrency: 2, d.ElementCryo: 1},
			cost:   map[d.Element]uint{d.ElementNone: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-2",
			origin: map[d.Element]uint{d.ElementCryo: 1, d.ElementCurrency: 1},
			cost:   map[d.Element]uint{d.ElementCurrency: 2},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-3",
			origin: map[d.Element]uint{d.ElementCryo: 1, d.ElementDendro: 1, d.ElementAnemo: 1},
			cost:   map[d.Element]uint{d.ElementCurrency: 3},
			want:   true,
		},
		{
			name:   "ElementSetEquals-Mixed-4",
			origin: map[d.Element]uint{d.ElementCryo: 2, d.ElementDendro: 2, d.ElementCurrency: 1},
			cost:   map[d.Element]uint{d.ElementCurrency: 5},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.ElementSetEqual(tt.origin, tt.cost); got != tt.want {
				t.Errorf("Error: %v, want %v", got, tt.want)
			}
		})
	}
}
