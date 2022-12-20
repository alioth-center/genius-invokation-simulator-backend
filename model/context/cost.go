package context

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type CostContext struct {
	need map[enum.ElementType]uint
}

// AddCost 增加指定ElementType的花费
func (c *CostContext) AddCost(element enum.ElementType, amount uint) {
	c.need[element] += amount
}

// SubCost 降低指定ElementType的花费
func (c *CostContext) SubCost(element enum.ElementType, amount uint) {
	if c.need[element] > amount {
		c.need[element] -= amount
	} else {
		c.need[element] = 0
	}
}

// Cost 返回CostContext携带的费用信息，只读
func (c *CostContext) Cost() map[enum.ElementType]uint {
	result := map[enum.ElementType]uint{}
	for elementType, cost := range c.need {
		result[elementType] = cost
	}

	return result
}

// NewCostContext 新建一个空CostContext
func NewCostContext() *CostContext {
	return &CostContext{
		need: map[enum.ElementType]uint{},
	}
}
