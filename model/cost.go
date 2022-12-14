/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "cost.go" LastUpdatedAt 2022/12/14 15:35:14
 */

package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/definition"
)

// MergeElementSet 将两个集合进行并集操作
func MergeElementSet(sets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	for _, set := range sets {
		for element, count := range set {
			result[element] += count
		}
	}

	return result
}

// MixElementSet 将两个集合进行交集操作
func MixElementSet(originSet definition.ElementSet, mixSets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	for element, count := range originSet {
		result[element] = count
	}

	for _, set := range mixSets {
		for element := definition.ElementStartIndex - 1; element <= definition.ElementEndIndex; element++ {
			if result[element] > set[element] {
				result[element] = set[element]
			}
		}
	}

	return result
}

// SubElementSet 将两个集合进行差集操作
func SubElementSet(originSet definition.ElementSet, subsets ...definition.ElementSet) (result definition.ElementSet) {
	result = map[definition.Element]uint{}
	for element, count := range originSet {
		result[element] = count
	}

	subset := MergeElementSet(subsets...)
	dropSet := MixElementSet(originSet, subset)
	for element, count := range dropSet {
		result[element] = result[element] - count
	}

	return result
}

// ElementSetStatistics 统计ElementSet的信息
func ElementSetStatistics(set definition.ElementSet) (elementCount, total uint) {
	for element := definition.ElementStartIndex - 1; element <= definition.ElementEndIndex; element++ {
		if count, ok := set[element]; ok && count != 0 {
			elementCount += 1
			total += count
		}
	}

	return elementCount, total
}

// ElementSetContains 判断subset是否为set的子集
func ElementSetContains(set, subset definition.ElementSet) bool {
	currencyElement, currencyElementCost, sameElementCost := subset[definition.ElementCurrency], subset[definition.ElementCurrency], subset[definition.ElementNone]
	delete(set, definition.ElementCurrency)
	delete(subset, definition.ElementCurrency)
	delete(subset, definition.ElementNone)

	for element := definition.ElementStartIndex; element <= definition.ElementEndIndex; element++ {
		if cost, ok := subset[element]; ok {
			if set[element]+currencyElement < cost {
				return false
			} else if set[element] < cost {
				currencyElement -= cost - subset[element]
			}
		}
	}

	leftSet := SubElementSet(set, subset)
	for element := definition.ElementStartIndex; element <= definition.ElementEndIndex; element++ {
		if _, totalCost := ElementSetStatistics(leftSet); leftSet[element]+currencyElement >= sameElementCost {
			return totalCost+currencyElement >= sameElementCost+currencyElementCost
		}
	}

	return false
}

// ElementSetEqual 判断originSet和costSet是否等价
func ElementSetEqual(originSet, costSet definition.ElementSet) bool {
	currencyElement, currencyElementCost, sameElementCost := originSet[definition.ElementCurrency], costSet[definition.ElementCurrency], costSet[definition.ElementNone]
	delete(originSet, definition.ElementCurrency)
	delete(costSet, definition.ElementCurrency)
	delete(costSet, definition.ElementNone)

	for element := definition.ElementStartIndex; element <= definition.ElementEndIndex; element++ {
		if cost, ok := costSet[element]; ok {
			if originSet[element]+currencyElement < cost {
				return false
			} else if originSet[element] < cost {
				currencyElement -= cost - originSet[element]
			}
		}
	}

	leftSet := SubElementSet(originSet, costSet)
	for element := definition.ElementStartIndex; element <= definition.ElementEndIndex; element++ {
		if _, totalCost := ElementSetStatistics(leftSet); leftSet[element]+currencyElement >= sameElementCost {
			return totalCost+currencyElement == sameElementCost+currencyElementCost
		}
	}

	return false
}

type CostContext struct {
	cost definition.ElementSet
}

func (c *CostContext) AddCosts(add definition.ElementSet) {
	c.cost = MergeElementSet(c.cost, add)
}

func (c *CostContext) SubCosts(sub definition.ElementSet) {
	c.cost = SubElementSet(c.cost, sub)
}

func (c *CostContext) EmptyCosts() {
	c.cost = map[definition.Element]uint{}
}
