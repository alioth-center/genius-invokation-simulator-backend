/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "element.go" LastUpdatedAt 2022/12/12 10:23:12
 */

package definition

// Element 元素类型，七元素与通用元素
type Element byte

// ElementSet 元素集合，用于回合开始的骰子
type ElementSet map[Element]uint

const (
	ElementCurrency  Element = iota     // ElementCurrency 通用元素
	ElementHydro                        // ElementHydro 水元素
	ElementPyro                         // ElementPyro 火元素
	ElementDendro                       // ElementDendro 草元素
	ElementGeo                          // ElementGeo 岩元素
	ElementCryo                         // ElementCryo 冰元素
	ElementElectro                      // ElementElectro 雷元素
	ElementAnemo                        // ElementAnemo 风元素
	ElementNone      Element = 1 << 4   // ElementNone 无元素
	ElementUndefined Element = 1<<8 - 1 // ElementUndefined 未定义元素，转化错误时传出此元素
)

// ToElement 将一个数转换为元素，如果是七元素与通用元素correct为true
func ToElement[T uint | int | uint64 | int64 | byte | rune](val T) (element Element, correct bool) {
	if val > T(ElementUndefined) || val < 0 {
		return ElementUndefined, false
	}

	result := Element(val)
	if result > ElementAnemo {
		return result, false
	} else {
		return result, true
	}
}
