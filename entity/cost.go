package entity

import (
	"math/rand"
	"time"

	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

var (
	random *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMilli()))
)

func init() {
	random.Seed(time.Now().UnixNano())
}

type Cost struct {
	costs map[enum.ElementType]uint
	total uint
}

// sub 从Cost中减少amount个element类型的元素骰子
func (c *Cost) sub(element enum.ElementType, amount uint) {
	if c.costs[element] >= amount {
		c.total -= amount
		c.costs[element] -= amount
	} else {
		c.total -= c.costs[element]
		c.costs[element] = 0
	}
}

// add 向Cost中增加amount个element类型的元素骰子
func (c *Cost) add(element enum.ElementType, amount uint) {
	c.total += amount
	c.costs[element] += amount
}

// Pay 从Cost中减去other中的元素骰子，不含等价判断
func (c *Cost) Pay(other Cost) {
	for element, amount := range other.costs {
		c.sub(element, amount)
	}
}

// Add 想Cost中增加other中的所有元素骰子，不含等价判断
func (c *Cost) Add(other Cost) {
	for element, amount := range other.costs {
		c.add(element, amount)
	}
}

// Contains 判断Cost中是否包含other中的所有元素，含等价判断
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

// Equals 判断Cost中的元素是否和other中的元素完全等价
func (c Cost) Equals(other Cost) bool {
	return c.Contains(other) && c.total == other.total
}

// NewCost 创建一个空Cost
func NewCost() *Cost {
	return &Cost{
		costs: map[enum.ElementType]uint{},
		total: 0,
	}
}

// NewCostFromMap 从一个map创建Cost
func NewCostFromMap(m map[enum.ElementType]uint) *Cost {
	result := &Cost{
		costs: map[enum.ElementType]uint{},
		total: 0,
	}

	for elementType, amount := range m {
		result.costs[elementType] = amount
		result.total += amount
	}

	return result
}

// NewRandomCost 创建一个随机费用的Cost
func NewRandomCost(elementAmount uint) *Cost {
	costMap := map[enum.ElementType]uint{}
	for elementAmount > 0 {
		if elementAmount <= 21 {
			source := random.Uint64()
			for i := uint(0); i < elementAmount; i++ {
				element := enum.ElementType(source % 8)
				costMap[element] += 1
				source = source >> 3
			}
			elementAmount = 0
		} else {
			source := random.Uint64()
			for i := 0; i < 21; i++ {
				element := enum.ElementType(source % 8)
				costMap[element] += 1
				source = source >> 3
			}
			elementAmount -= 21
		}
	}
	return NewCostFromMap(costMap)
}
