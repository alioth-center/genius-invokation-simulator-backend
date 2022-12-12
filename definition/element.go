/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "element.go" LastUpdatedAt 2022/12/12 10:23:12
 */

package definition

import "errors"

// Element 元素类型，七元素与通用元素
type Element byte

// ElementSet 元素集合，用于回合开始的骰子
type ElementSet map[Element]uint

const (
	ElementCurrency Element = iota // ElementCurrency 通用元素
	ElementHydro                   // ElementHydro 水元素
	ElementPyro                    // ElementPyro 火元素
	ElementDendro                  // ElementDendro 草元素
	ElementGeo                     // ElementGeo 岩元素
	ElementCryo                    // ElementCryo 冰元素
	ElementElectro                 // ElementElectro 雷元素
	ElementAnemo                   // ElementAnemo 风元素
)

// ToElement 将一个数转换为元素，如果是七元素与通用元素correct为true
func ToElement[T uint | int | uint64 | int64 | byte | rune](val T) (element Element, correct bool) {
	if val > 255 || val < 0 {
		return 255, false
	}

	result := Element(val)
	if result > ElementAnemo {
		return result, false
	} else {
		return result, true
	}
}

// ToElementSet 用一个随机数随机数生成元素集合，三个二进制位生成一个元素
func ToElementSet[T uint | int | uint64 | int64](randomNumber T, elementCount uint) (result ElementSet) {
	result = map[Element]uint{}
	for i := uint(0); i < elementCount; i++ {
		element, _ := ToElement(randomNumber % 8)
		result[element] += 1
		randomNumber = randomNumber >> 3
	}

	return result
}

// MergeElementSet 将若干个元素集合进行并集运算
func MergeElementSet(elementSets ...ElementSet) (result ElementSet) {
	result = map[Element]uint{}
	for _, set := range elementSets {
		for element, count := range set {
			result[element] += count
		}
	}

	return result
}

// SubElementSet 将若干个元素集合进行差集运算，第一个集合是被操作的集合，参数小于两个的时候会报错
func SubElementSet(elementSets ...ElementSet) (result ElementSet, err error) {
	if len(elementSets) < 2 {
		return nil, errors.New("incorrect elementSets")
	} else {
		result = elementSets[0]
		for _, set := range elementSets[1:] {
			for element, count := range set {
				if number, ok := set[element]; !ok || number-count < 0 {
					return nil, errors.New("incorrect subSet")
				} else {
					result[element] = number - count
				}
			}
		}
		return result, nil
	}
}
