package persistence

import "github.com/sunist-c/genius-invokation-simulator-backend/entity"

type CharacterInfo struct {
	PersistenceIndex[entity.CharacterInfo]
	Skills []uint
}
