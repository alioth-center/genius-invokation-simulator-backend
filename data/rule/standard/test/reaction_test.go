/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "reaction_test.go" LastUpdatedAt 2022/12/13 11:07:13
 */

package test

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/data/rule/standard"
	d "github.com/sunist-c/genius-invokation-simulator-backend/definition"

	"reflect"
	"testing"
)

func TestReactionCalculatorImplement_Calculate(t *testing.T) {
	type args struct {
		elementNew      d.Element
		elementAttached []d.Element
	}
	tests := []struct {
		name               string
		args               args
		wantReaction       d.Reaction
		wantElementSurplus []d.Element
	}{
		{
			name: "ReactionCalculator-NoElement-Attachable",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementElectro, elementAttached: []d.Element{}},
			wantReaction:       d.ReactionNone,
			wantElementSurplus: []d.Element{d.ElementElectro},
		},
		{
			name: "ReactionCalculator-NoElement-NonAttachable",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementGeo, elementAttached: []d.Element{}},
			wantReaction:       d.ReactionNone,
			wantElementSurplus: []d.Element{},
		},
		{
			name: "ReactionCalculator-OneElement-Reactive",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementElectro, elementAttached: []d.Element{d.ElementPyro}},
			wantReaction:       d.ReactionOverloaded,
			wantElementSurplus: []d.Element{},
		},
		{
			name: "ReactionCalculator-OneElement-Unreactive-Attachable",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementDendro, elementAttached: []d.Element{d.ElementCryo}},
			wantReaction:       d.ReactionNone,
			wantElementSurplus: []d.Element{d.ElementCryo, d.ElementDendro},
		},
		{
			name: "ReactionCalculator-OneElement-Unreactive-NonAttachable",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementAnemo, elementAttached: []d.Element{d.ElementDendro}},
			wantReaction:       d.ReactionNone,
			wantElementSurplus: []d.Element{d.ElementDendro},
		},
		{
			name: "ReactionCalculator-OneElement-Unreactive-SameElement",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementPyro, elementAttached: []d.Element{d.ElementPyro}},
			wantReaction:       d.ReactionNone,
			wantElementSurplus: []d.Element{d.ElementPyro},
		},
		{
			name: "ReactionCalculator-TwoElement-Reactive-1",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementPyro, elementAttached: []d.Element{d.ElementDendro, d.ElementCryo}},
			wantReaction:       d.ReactionBurning,
			wantElementSurplus: []d.Element{d.ElementCryo},
		},
		{
			name: "ReactionCalculator-TwoElement-Reactive-2",
			args: struct {
				elementNew      d.Element
				elementAttached []d.Element
			}{elementNew: d.ElementPyro, elementAttached: []d.Element{d.ElementCryo, d.ElementDendro}},
			wantReaction:       d.ReactionMelt,
			wantElementSurplus: []d.Element{d.ElementDendro},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := standard.ReactionCalculatorImplement{}
			gotReaction, gotElementSurplus := r.Calculate(tt.args.elementNew, tt.args.elementAttached)
			if gotReaction != tt.wantReaction {
				t.Errorf("Calculate() gotReaction = %v, want %v", gotReaction, tt.wantReaction)
			}
			if !reflect.DeepEqual(gotElementSurplus, tt.wantElementSurplus) {
				t.Errorf("Calculate() gotElementSurplus = %v, want %v", gotElementSurplus, tt.wantElementSurplus)
			}
		})
	}
}

func TestReactionCalculatorImplement_Type(t *testing.T) {
	tests := []struct {
		name string
		want d.RuleType
	}{
		{name: "ReactionCalculator-Type", want: d.RuleInGameModifier},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := standard.ReactionCalculatorImplement{}
			if got := r.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
