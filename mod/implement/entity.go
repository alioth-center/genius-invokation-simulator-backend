package implement

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
)

type EntityImpl struct {
	typeID uint64
}

func (b *EntityImpl) TypeID() uint64 {
	return b.typeID
}

func (b *EntityImpl) InjectTypeID(id uint64) {
	b.typeID = util.GenerateRealID(ModID(), uint16(id))
}

type EntityOptions func(entity *EntityImpl)

func WithEntityID(id uint16) EntityOptions {
	return func(entity *EntityImpl) {
		entity.InjectTypeID(uint64(id))
	}
}

func NewEntityWithOpts(options ...EntityOptions) model.BaseEntity {
	entity := &EntityImpl{}
	for _, o := range options {
		o(entity)
	}

	return entity
}
