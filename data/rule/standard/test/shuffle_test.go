/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "shuffle_test.go" LastUpdatedAt 2022/12/13 10:53:13
 */

package test

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/data/rule/standard"
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"

	"fmt"
	"testing"
)

type testCardImpl struct {
	id uint
}

func (t testCardImpl) ID() uint {
	return t.id
}

func (t testCardImpl) Name() string {
	panic("implement me")
}

func (t testCardImpl) Description() string {
	panic("implement me")
}

func (t testCardImpl) Type() definition.CardType {
	panic("implement me")
}

func (t testCardImpl) Cost() definition.ElementSet {
	panic("implement me")
}

func newTestCards(count uint) []model.ICard {
	result := make([]model.ICard, count)
	for i := uint(0); i < count; i++ {
		result[i] = testCardImpl{id: i}
	}

	return result
}

func newPlayers(count uint) []*model.Player {
	result := make([]*model.Player, count)
	for i := uint(0); i < count; i++ {
		result[i] = &model.Player{UID: i}
	}

	return result
}

func TestShuffleCardStackImplement_Shuffle(t *testing.T) {
	type args struct {
		start int
		end   int
		array []model.ICard
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "ShuffleCardStack-1", args: struct {
			start int
			end   int
			array []model.ICard
		}{start: 0, end: 4, array: newTestCards(4)}},
		{name: "ShuffleCardStack-2", args: struct {
			start int
			end   int
			array []model.ICard
		}{start: 0, end: 10, array: newTestCards(10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := standard.NewShuffleCardStackImplement()
			s.Shuffle(tt.args.start, tt.args.end, tt.args.array)
			fmt.Println(tt.args.array)
		})
	}
}

func TestShuffleCardStackImplement_Type(t *testing.T) {
	tests := []struct {
		name string
		want definition.RuleType
	}{
		{name: "ShuffleCardStack-Type", want: definition.RuleInitializeGame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := standard.NewShuffleCardStackImplement()
			if got := s.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShufflePlayerChainImplement_Shuffle(t *testing.T) {
	type args struct {
		start int
		end   int
		array []*model.Player
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "ShufflePlayerChain-1", args: struct {
			start int
			end   int
			array []*model.Player
		}{start: 0, end: 4, array: newPlayers(4)}},
		{name: "ShufflePlayerChain-2", args: struct {
			start int
			end   int
			array []*model.Player
		}{start: 0, end: 10, array: newPlayers(10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := standard.NewShufflePlayerChainImplement()
			s.Shuffle(tt.args.start, tt.args.end, tt.args.array)
			fmt.Printf("%+v\n", tt.args.array)
		})
	}
}

func TestShufflePlayerChainImplement_Type(t *testing.T) {
	tests := []struct {
		name string
		want definition.RuleType
	}{
		{name: "ShufflePlayerChain-Type", want: definition.RuleInitializeGame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := standard.NewShufflePlayerChainImplement()
			if got := s.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
