package context

type ChargeContext struct {
	charges map[uint]int
}

// AddMagic 向指定目标添加能量
func (c *ChargeContext) AddMagic(target, amount uint) {
	c.charges[target] += int(amount)
}

// SubMagic 减少指定目标的能量
func (c *ChargeContext) SubMagic(target, amount uint) {
	c.charges[target] -= int(amount)
}

// Charge 返回ChargeContext携带的充能信息，只读
func (c ChargeContext) Charge() map[uint]int {
	result := map[uint]int{}
	for target, amount := range c.charges {
		result[target] = amount
	}

	return result
}
