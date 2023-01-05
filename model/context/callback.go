package context

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type CallbackContext struct {
	changeElements  kv.Pair[bool, *CostContext]
	changeCharge    kv.Pair[bool, *ChargeContext]
	changeModifiers kv.Pair[bool, *ModifierContext]
	attachElement   kv.Pair[bool, map[uint]enum.ElementType]
	getCards        kv.Pair[bool, uint]
	findCard        kv.Pair[bool, enum.CardType]
	switchCharacter kv.Pair[bool, uint]
	operated        kv.Pair[bool, bool]
}

func (c *CallbackContext) ChangeElements(f func(ctx *CostContext)) {
	if !c.changeElements.Key() {
		c.changeElements.SetKey(true)
	}
	f(c.changeElements.Value())
}

func (c *CallbackContext) ChangeCharge(f func(ctx *ChargeContext)) {
	if !c.changeCharge.Key() {
		c.changeCharge.SetKey(true)
	}
	f(c.changeCharge.Value())
}

func (c *CallbackContext) ChangeModifiers(f func(ctx *ModifierContext)) {
	if !c.changeModifiers.Key() {
		c.changeModifiers.SetKey(true)
	}
	f(c.changeModifiers.Value())
}

func (c *CallbackContext) AttachElement(target uint, element enum.ElementType) {
	if !c.attachElement.Key() {
		c.attachElement.SetKey(true)
	}
	value := c.attachElement.Value()
	value[target] = element
	c.attachElement.SetValue(value)
}

func (c *CallbackContext) GetCards(amount uint) {
	if !c.getCards.Key() {
		c.getCards.SetKey(true)
	}
	c.getCards.SetValue(amount)
}

func (c *CallbackContext) FindCard(cardType enum.CardType) {
	if !c.findCard.Key() {
		c.findCard.SetKey(true)
	}
	c.findCard.SetValue(cardType)
}

func (c *CallbackContext) SwitchCharacter(target uint) {
	if !c.switchCharacter.Key() {
		c.switchCharacter.SetKey(true)
	}
	c.switchCharacter.SetValue(target)
}

func (c *CallbackContext) ChangeOperated(operated bool) {
	if !c.operated.Key() {
		c.operated.SetKey(true)
	}
	c.operated.SetValue(operated)
}

func (c CallbackContext) ChangeElementsResult() (changed bool, result *CostContext) {
	return c.changeElements.Key(), c.changeElements.Value()
}

func (c CallbackContext) ChangeChargeResult() (changed bool, result *ChargeContext) {
	return c.changeCharge.Key(), c.changeCharge.Value()
}

func (c CallbackContext) ChangeModifiersResult() (changed bool, result *ModifierContext) {
	return c.changeModifiers.Key(), c.changeModifiers.Value()
}

func (c CallbackContext) AttachElementResult() (changed bool, result map[uint]enum.ElementType) {
	return c.attachElement.Key(), c.attachElement.Value()
}

func (c CallbackContext) GetCardsResult() (changed bool, result uint) {
	return c.getCards.Key(), c.getCards.Value()
}

func (c CallbackContext) GetFindCardResult() (find bool, cardType enum.CardType) {
	return c.findCard.Key(), c.findCard.Value()
}

func (c CallbackContext) SwitchCharacterResult() (switched bool, target uint) {
	return c.switchCharacter.Key(), c.switchCharacter.Value()
}

func (c CallbackContext) ChangeOperatedResult() (switched, operated bool) {
	return c.operated.Key(), c.operated.Value()
}

func NewCallbackContext() *CallbackContext {
	return &CallbackContext{
		changeElements:  kv.NewPair(false, NewCostContext()),
		changeCharge:    kv.NewPair(false, NewChargeContext()),
		changeModifiers: kv.NewPair(false, NewModifierContext()),
		attachElement:   kv.NewPair(false, map[uint]enum.ElementType{}),
		getCards:        kv.NewPair(false, uint(0)),
		switchCharacter: kv.NewPair(false, uint(0)),
		operated:        kv.NewPair(false, false),
		findCard:        kv.NewPair(false, enum.CardType(0)),
	}
}
