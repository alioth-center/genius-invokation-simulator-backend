package persistence

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type Card struct {
	PersistenceIndex[entity.Card]
	Type enum.CardType
}
