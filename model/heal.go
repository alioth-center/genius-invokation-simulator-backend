/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "heal.go" LastUpdatedAt 2022/12/15 13:26:15
 */

package model

type HealContext struct {
	amount uint
}

func (c *HealContext) AddHeal(amount uint) {
	c.amount += amount
}

func (c *HealContext) SubHeal(amount uint) {
	if c.amount > amount {
		c.amount -= amount
	} else {
		c.amount = 0
	}
}

func (c *HealContext) CancelHeal() {
	c.amount = 0
}

func (c HealContext) Heal() uint {
	return c.amount
}

func NewHealContext(heal uint) *HealContext {
	return &HealContext{amount: heal}
}
