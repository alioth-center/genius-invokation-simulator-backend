package context

type ChargeContext struct {
	charges map[uint64]int
}

// AddMagic 向指定目标添加能量
func (c *ChargeContext) AddMagic(target uint64, amount uint) {
	c.charges[target] += int(amount)
}

// SubMagic 减少指定目标的能量
func (c *ChargeContext) SubMagic(target uint64, amount uint) {
	c.charges[target] -= int(amount)
}

// Charge 返回ChargeContext携带的充能信息，只读
func (c ChargeContext) Charge() map[uint64]int {
	result := map[uint64]int{}
	for target, amount := range c.charges {
		result[target] = amount
	}

	return result
}

// NewChargeContext 新建一个空的ChargeContext
func NewChargeContext() *ChargeContext {
	return &ChargeContext{
		charges: map[uint64]int{},
	}
}
