package enum

// ElementType 元素类型，七元素与通用元素
type ElementType byte

const (
	ElementCurrency  ElementType = iota     // ElementCurrency 通用元素，用于表示任意元素
	ElementAnemo                            // ElementAnemo 风元素
	ElementCryo                             // ElementCryo 冰元素
	ElementDendro                           // ElementDendro 草元素
	ElementElectro                          // ElementElectro 雷元素
	ElementGeo                              // ElementGeo 岩元素
	ElementHydro                            // ElementHydro 水元素
	ElementPyro                             // ElementPyro 火元素
	ElementNone      ElementType = 1 << 4   // ElementNone 无元素，用于表示物理攻击
	ElementSame      ElementType = 1 << 4   // ElementNone 无元素，用于表示相同元素
	ElementUndefined ElementType = 1<<8 - 1 // ElementUndefined 未定义元素，转化错误时传出此元素

	ElementStartIndex ElementType = 1 // ElementStartIndex 七元素起始
	ElementEndIndex   ElementType = 7 // ElementEndIndex 七元素终止
)
