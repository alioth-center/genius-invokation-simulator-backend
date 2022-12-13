/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "shuffle.go" LastUpdatedAt 2022/12/12 15:38:12
 */

package standard

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"

	"math/rand"
	"time"
)

type ShuffleCardStackImplement struct {
	random      *rand.Rand
	initialized bool
}

func shuffle[T any](random *rand.Rand, start, end int, array []T) {
	random.Shuffle(len(array[start:end]), func(i, j int) {
		array[i+start], array[j+start] = array[j+start], array[i+start]
	})
}

func (s *ShuffleCardStackImplement) Shuffle(start, end int, array []model.ICard) {
	if s.initialized {
		shuffle(s.random, start, end, array)
	} else {
		s.random.Seed(time.Now().UnixNano())
		s.initialized = true
		s.Shuffle(start, end, array)
	}
}

func (s ShuffleCardStackImplement) Type() definition.RuleType {
	return definition.RuleInitializeGame
}

type ShufflePlayerChainImplement struct {
	random      *rand.Rand
	initialized bool
}

func (s ShufflePlayerChainImplement) Type() definition.RuleType {
	return definition.RuleInitializeGame
}

func (s *ShufflePlayerChainImplement) Shuffle(start, end int, array []*model.Player) {
	if s.initialized {
		shuffle(s.random, start, end, array)
	} else {
		s.random.Seed(time.Now().UnixNano())
		s.initialized = true
		s.Shuffle(start, end, array)
	}
}

func NewShuffleCardStackImplement() model.EventShuffleCardStackInterface {
	impl := &ShuffleCardStackImplement{}
	impl.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	impl.initialized = true
	return impl
}

func NewShufflePlayerChainImplement() model.EventShufflePlayerChainInterface {
	impl := &ShufflePlayerChainImplement{}
	impl.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	impl.initialized = true
	return impl
}
