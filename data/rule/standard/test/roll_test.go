/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "roll_test.go" LastUpdatedAt 2022/12/13 13:25:13
 */

package test

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/data/rule/standard"
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"testing"
)

func TestRollStageHandlerImplement_ReRoll(t *testing.T) {
	type args struct {
		originSet definition.ElementSet
		dropSet   definition.ElementSet
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "RollStageHandler-ReRoll-1", args: args{
			originSet: map[definition.Element]uint{definition.ElementCurrency: 8},
			dropSet:   map[definition.Element]uint{definition.ElementCurrency: 4},
		}},
		{name: "RollStageHandler-ReRoll-2", args: args{
			originSet: map[definition.Element]uint{definition.ElementCurrency: 4},
			dropSet:   map[definition.Element]uint{definition.ElementCurrency: 6},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := standard.NewRollStageHandlerImplement()
			got := r.ReRoll(tt.args.originSet, tt.args.dropSet)
			fmt.Printf("ReRoll Result: %+v\n", got)
		})
	}
}

func TestRollStageHandlerImplement_Roll(t *testing.T) {
	type args struct {
		setCaps uint
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "RollStageHandler-Roll-1", args: args{setCaps: 8}},
		{name: "RollStageHandler-Roll-2", args: args{setCaps: 4}},
		{name: "RollStageHandler-Roll-3", args: args{setCaps: 1}},
		{name: "RollStageHandler-Roll-4", args: args{setCaps: 0}},
		{name: "RollStageHandler-Roll-5", args: args{setCaps: 114514}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := standard.NewRollStageHandlerImplement()
			got := r.Roll(tt.args.setCaps)
			fmt.Printf("%+v\n", got)
			gotCount := uint(0)
			for _, v := range got {
				gotCount += v
			}
			if gotCount != tt.args.setCaps {
				t.Errorf("expected element-set size, want %v, got %v", tt.args.setCaps, len(got))
			}
		})
	}
}

func TestRollStageHandlerImplement_Type(t *testing.T) {
	tests := []struct {
		name string
		want definition.RuleType
	}{
		{name: "RollStageHandler-Type", want: definition.RuleInGameModifier},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := standard.NewRollStageHandlerImplement()
			if got := r.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
