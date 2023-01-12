package model

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
)

type Card interface {
	ID() uint
	Type() enum.CardType
}

type FoodCard interface {
	Card
	ExecuteModify(ctx *context.ModifierContext)
}
