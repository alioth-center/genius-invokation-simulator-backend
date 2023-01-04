package context

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
)

type CallbackContext struct {
	changeElements  *CostContext
	changeCharge    *ChargeContext
	changeModifiers *ModifierContext
	attachElement   map[uint]enum.ElementType
	getCards        uint
	findCard        kv.Pair[bool, enum.CardType]
	switchCharacter kv.Pair[bool, uint]
	operated        kv.Pair[bool, bool]
}

func (c *CallbackContext) ChangeElements(f func(ctx *CostContext)) {
	f(c.changeElements)
}

func (c *CallbackContext) ChangeCharge(f func(ctx *ChargeContext)) {
	f(c.changeCharge)
}

func (c *CallbackContext) ChangeModifiers(f func(ctx *ModifierContext)) {
	f(c.changeModifiers)
}

func (c *CallbackContext) AttachElement(target uint, element enum.ElementType) {
	c.attachElement[target] = element
}

func (c *CallbackContext) GetCards(amount uint) {
	c.getCards = amount
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

func (c CallbackContext) ChangeElementsResult() CostContext {
	return *c.changeElements
}

func (c CallbackContext) ChangeChargeResult() ChargeContext {
	return *c.changeCharge
}

func (c CallbackContext) ChangeModifiersResult() ModifierContext {
	return *c.changeModifiers
}

func (c CallbackContext) AttachElementResult() map[uint]enum.ElementType {
	return c.attachElement
}

func (c CallbackContext) GetCardsResult() uint {
	return c.getCards
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
		changeElements:  NewCostContext(),
		changeCharge:    NewChargeContext(),
		changeModifiers: NewModifierContext(),
		attachElement:   map[uint]enum.ElementType{},
		getCards:        0,
		switchCharacter: kv.NewPair(false, uint(0)),
		operated:        kv.NewPair(false, false),
	}
}
