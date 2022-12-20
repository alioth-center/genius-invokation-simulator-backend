package entity

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type Cost struct {
	costs map[enum.ElementType]uint
	total uint
}

func (c Cost) Contains(other Cost) bool {
	// 判断费用总数
	if other.total > c.total {
		return false
	}

	// 先减去确定类型的费用
	for element := enum.ElementStartIndex; element <= enum.ElementEndIndex; element++ {
		if other.costs[element] > c.costs[element]+c.costs[enum.ElementCurrency] {
			return false
		} else {
			if other.costs[element] > c.costs[element] {
				c.costs[enum.ElementCurrency] -= other.costs[element] - c.costs[element]
				c.costs[element] = 0
			} else {
				c.costs[element] -= other.costs[element]
			}
		}
	}

	// 判断是否满足同色元素要求
	for element := enum.ElementStartIndex; element <= enum.ElementEndIndex; element++ {
		if other.costs[enum.ElementSame] <= c.costs[element]+c.costs[enum.ElementCurrency] && element != enum.ElementCurrency {
			return true
		}
	}

	return false
}

func (c Cost) Equals(other Cost) bool {
	return c.Contains(other) && c.total == other.total
}

func newCostFromMap(m map[enum.ElementType]uint) *Cost {
	result := &Cost{
		costs: map[enum.ElementType]uint{},
	}
	for elementType, amount := range m {
		result.costs[elementType] = amount
		result.total += amount
	}
	return result
}
