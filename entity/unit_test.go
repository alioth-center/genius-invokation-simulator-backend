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
			originCost := NewCostFromMap(tt.origin)
			otherCost := NewCostFromMap(tt.cost)
			if result := originCost.Equals(*otherCost); result != tt.want {
				t.Errorf("incorrect result, want %v, got %v", tt.want, result)
			}
		})
	}
}

func TestCostAdd(t *testing.T) {
	tests := []struct {
		name   string
		origin map[enum.ElementType]uint
		other  map[enum.ElementType]uint
		want   map[enum.ElementType]uint
	}{
		{
			name:   "TestCostAdd-1",
			origin: map[enum.ElementType]uint{},
			other:  map[enum.ElementType]uint{enum.ElementCryo: 114514},
			want:   map[enum.ElementType]uint{enum.ElementCryo: 114514},
		},
		{
			name:   "TestCostAdd-2",
			origin: map[enum.ElementType]uint{enum.ElementPyro: 1},
			other:  map[enum.ElementType]uint{enum.ElementCryo: 1},
			want:   map[enum.ElementType]uint{enum.ElementCryo: 1, enum.ElementPyro: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cost := NewCostFromMap(tt.origin)
			other := NewCostFromMap(tt.other)
			cost.Add(*other)
			for element, amount := range cost.costs {
				if tt.want[element] != amount {
					t.Errorf("incorrect result, want %v, got %v", tt.want[element], amount)
				}
			}
		})
	}
}

func TestCostSub(t *testing.T) {
	tests := []struct {
		name   string
		origin map[enum.ElementType]uint
		other  map[enum.ElementType]uint
		want   map[enum.ElementType]uint
	}{
		{
			name:   "TestCostSub-1",
			origin: map[enum.ElementType]uint{enum.ElementCryo: 114514},
			other:  map[enum.ElementType]uint{enum.ElementCryo: 514},
			want:   map[enum.ElementType]uint{enum.ElementCryo: 114000},
		},
		{
			name:   "TestCostSub-2",
			origin: map[enum.ElementType]uint{enum.ElementPyro: 1},
			other:  map[enum.ElementType]uint{enum.ElementCryo: 1},
			want:   map[enum.ElementType]uint{enum.ElementPyro: 1, enum.ElementCryo: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cost := NewCostFromMap(tt.origin)
			other := NewCostFromMap(tt.other)
			cost.Pay(*other)
			for element, amount := range cost.costs {
				if tt.want[element] != amount {
					t.Errorf("incorrect result, want %v, got %v", tt.want[element], amount)
				}
			}
		})
	}
}

type testCard struct {
	id uint
	t  enum.CardType
}

func (t testCard) ID() uint { return t.id }

func (t testCard) Type() enum.CardType { return t.t }

func newTestCard(id uint, t enum.CardType) Card { return testCard{id: id, t: t} }

func TestCardDeckGet(t *testing.T) {
	food := newTestCard(1, enum.CardFood)
	companion := newTestCard(2, enum.CardCompanion)
	item := newTestCard(3, enum.CardItem)
	tests := []struct {
		name       string
		cards      []Card
		got        int
		wantCard   Card
		wantResult bool
	}{
		{
			name:       "TestCardDeckGet-1",
			cards:      []Card{food, companion, item},
			got:        2,
			wantCard:   item,
			wantResult: true,
		},
		{
			name:       "TestCardDeckGet-2",
			cards:      []Card{food, companion, item},
			got:        3,
			wantCard:   nil,
			wantResult: false,
		},
		{
			name:       "TestCardDeckGet-3",
			cards:      []Card{food, companion, item},
			got:        0,
			wantCard:   food,
			wantResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewCardDeck(tt.cards)
			for i := 0; i < tt.got; i++ {
				deck.GetOne()
			}
			if card, success := deck.GetOne(); success != tt.wantResult {
				t.Errorf("got %v, want %v", success, tt.wantResult)
			} else if card != tt.wantCard {
				t.Errorf("got %+v, want %+v", card, tt.wantCard)
			}
		})
	}
}

func TestCardDeckFind(t *testing.T) {
	food := newTestCard(1, enum.CardFood)
	companion := newTestCard(2, enum.CardCompanion)
	item := newTestCard(3, enum.CardItem)
	tests := []struct {
		name       string
		cards      []Card
		got        int
		find       enum.CardType
		wantCard   Card
		wantResult bool
	}{
		{
			name:       "TestCardDeckFind-1",
			cards:      []Card{food, companion, item},
			got:        2,
			find:       enum.CardFood,
			wantCard:   nil,
			wantResult: false,
		},
		{
			name:       "TestCardDeckFind-2",
			cards:      []Card{food, companion, item},
			got:        0,
			find:       enum.CardArtifact,
			wantCard:   nil,
			wantResult: false,
		},
		{
			name:       "TestCardDeckFind-3",
			cards:      []Card{food, companion, item},
			got:        0,
			find:       enum.CardItem,
			wantCard:   item,
			wantResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewCardDeck(tt.cards)
			for i := 0; i < tt.got; i++ {
				deck.GetOne()
			}
			if card, success := deck.FindOne(tt.find); success != tt.wantResult {
				t.Errorf("got %v, want %v", success, tt.wantResult)
			} else if card != tt.wantCard {
				t.Errorf("got %+v, want %+v", card, tt.wantCard)
			}
		})
	}
}

func TestCardDeckReset(t *testing.T) {
	food := newTestCard(1, enum.CardFood)
	companion := newTestCard(2, enum.CardCompanion)
	item := newTestCard(3, enum.CardItem)
	tests := []struct {
		name       string
		cards      []Card
		got        int
		holdings   []uint
		wantCard   Card
		wantResult bool
	}{
		{
			name:       "TestCardDeckReset-1",
			cards:      []Card{food, companion, item},
			got:        3,
			holdings:   []uint{},
			wantCard:   food,
			wantResult: true,
		},
		{
			name:       "TestCardDeckReset-2",
			cards:      []Card{food, companion, item},
			got:        3,
			holdings:   []uint{1, 2, 3},
			wantCard:   nil,
			wantResult: false,
		},
		{
			name:       "TestCardDeckReset-3",
			cards:      []Card{food, companion, item},
			got:        0,
			holdings:   []uint{1, 2},
			wantCard:   item,
			wantResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewCardDeck(tt.cards)
			for i := 0; i < tt.got; i++ {
				deck.GetOne()
			}
			deck.Reset(tt.holdings)
			if card, success := deck.GetOne(); success != tt.wantResult {
				t.Errorf("got %v, want %v", success, tt.wantResult)
			} else if card != tt.wantCard {
				t.Errorf("got %+v, want %+v", card, tt.wantCard)
			}
		})
	}
}

func TestCardDeckShuffle(t *testing.T) {
	food := newTestCard(1, enum.CardFood)
	companion := newTestCard(2, enum.CardCompanion)
	item := newTestCard(3, enum.CardItem)
	tests := []struct {
		name       string
		cards      []Card
		got        int
		wontCard   Card
		wontResult bool
	}{
		{
			name:       "TestCardDeckShuffle-1",
			cards:      []Card{food, companion, item},
			got:        1,
			wontCard:   food,
			wontResult: false,
		},
		{
			name:       "TestCardDeckShuffle-2",
			cards:      []Card{food, companion, item},
			got:        0,
			wontCard:   nil,
			wontResult: false,
		},
		{
			name:       "TestCardDeckShuffle-3",
			cards:      []Card{food, companion, item},
			got:        3,
			wontCard:   food,
			wontResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cardDeck := NewCardDeck(tt.cards)
			for i := 0; i < tt.got; i++ {
				cardDeck.GetOne()
			}

			cardDeck.Shuffle()

			if card, success := cardDeck.GetOne(); success == tt.wontResult {
				t.Errorf("got %v, wont %v", success, tt.wontResult)
			} else if card == tt.wontCard {
				t.Errorf("got %+v, wont %+v", card, tt.wontCard)
			}
		})
	}
}
