/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "shuffle.go" LastUpdatedAt 2022/12/12 15:38:12
 */

package standard

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"

	"math/rand"
)

type ShuffleCardStackImplement struct{}

func shuffle[T any](start, end int, array []T) {
	rand.Shuffle(len(array[start:end]), func(i, j int) {
		array[i+start], array[j+start] = array[j+start], array[i+start]
	})
}

func (s ShuffleCardStackImplement) Shuffle(start, end int, array []model.ICard) {
	shuffle(start, end, array)
}

func (s ShuffleCardStackImplement) Type() definition.RuleType {
	return definition.RuleInitializeGame
}

type ShufflePlayerChainImplement struct{}

func (s ShufflePlayerChainImplement) Type() definition.RuleType {
	return definition.RuleInitializeGame
}

func (s ShufflePlayerChainImplement) Shuffle(start, end int, array []*model.Player) {
	shuffle(start, end, array)
}
