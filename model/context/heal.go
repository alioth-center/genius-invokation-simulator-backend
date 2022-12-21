package context

type HealContext struct {
	heals map[uint]uint
}

// AddHeal 向指定目标添加治疗量
func (h *HealContext) AddHeal(target, amount uint) {
	h.heals[target] += amount
}

// SubHeal 减少指定目标的治疗量
func (h *HealContext) SubHeal(target, amount uint) {
	if h.heals[target] > amount {
		h.heals[target] -= amount
	} else {
		h.heals[target] = 0
	}
}

// Heal 返回HealContext携带的治疗信息，只读
func (h HealContext) Heal() map[uint]uint {
	result := map[uint]uint{}
	for target, amount := range h.heals {
		result[target] = amount
	}

	return result
}

// NewHealContext 新建一个空的HealContext
func NewHealContext() *HealContext {
	return &HealContext{
		heals: map[uint]uint{},
	}
}
