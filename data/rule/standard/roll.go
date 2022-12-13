/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "roll.go" LastUpdatedAt 2022/12/13 10:43:13
 */

package standard

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
	"github.com/sunist-c/genius-invokation-simulator-backend/model"
	"math/rand"
	"time"
)

type RollStageHandlerImplement struct {
	random      *rand.Rand
	initialized bool
}

func (r *RollStageHandlerImplement) generateRandomNumbers(caps uint) (numbers []uint64) {
	if r.initialized {
		numbers := make([]uint64, caps)
		for i := uint(0); i < caps; i++ {
			numbers[i] = r.random.Uint64()
		}
		return numbers
	} else {
		r.random.Seed(time.Now().UnixNano())
		r.initialized = true
		return r.generateRandomNumbers(caps)
	}
}

func generateElementSet(count uint, random uint64) (generated bool, elements definition.ElementSet) {
	if count < 22 {
		elements = map[definition.Element]uint{}
		for i := uint(0); i < count; i++ {
			element, _ := definition.ToElement(random % 8)
			elements[element] += 1
			random = random >> 3
		}
		return true, elements
	} else {
		return false, nil
	}
}

func mergeElementSet(sets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	for _, set := range sets {
		for element, count := range set {
			result[element] += count
		}
	}

	return result
}

func mixElementSet(originSet definition.ElementSet, mixSets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	for element, count := range originSet {
		result[element] = count
	}

	for _, set := range mixSets {
		for element := definition.Element(0); element <= definition.ElementAnemo; element++ {
			if _, ok := originSet[element]; ok {
				if result[element] > set[element] {
					result[element] = set[element]
				}
			}
		}
	}

	return result
}

func subElementSet(originSet definition.ElementSet, subsets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	subset := mergeElementSet(subsets...)
	dropSet := mixElementSet(originSet, subset)
	for element, count := range dropSet {
		result[element] = originSet[element] - count
	}

	return result
}

func (r RollStageHandlerImplement) Type() definition.RuleType {
	return definition.RuleInGameModifier
}

func (r RollStageHandlerImplement) Roll(setCaps uint) (set definition.ElementSet) {
	set = map[definition.Element]uint{}
	randomNumCount := setCaps/21 + 1
	randomNumbers := r.generateRandomNumbers(randomNumCount)
	subsets := make([]definition.ElementSet, randomNumCount)
	generated := uint(0)
	for i, random := range randomNumbers {
		if setCaps-generated <= 21 {
			_, subset := generateElementSet(setCaps-generated, random)
			subsets[i] = subset
			break
		} else {
			_, subset := generateElementSet(21, random)
			subsets[i] = subset
			generated += 21
		}
	}

	return mergeElementSet(subsets...)
}

func (r RollStageHandlerImplement) ReRoll(originSet definition.ElementSet, dropSet definition.ElementSet) (result definition.ElementSet) {
	canDropSet := mixElementSet(originSet, dropSet)
	reRollNumber := uint(0)
	for _, count := range canDropSet {
		reRollNumber += count
	}

	reRolledSet := r.Roll(reRollNumber)

	return mergeElementSet(subElementSet(originSet, canDropSet), reRolledSet)
}

func NewRollStageHandlerImplement() model.EventRollStageHandlerInterface {
	impl := &RollStageHandlerImplement{}
	impl.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	impl.initialized = true
	return impl
}
