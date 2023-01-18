package implement

type BaseEntityImpl struct {
	typeID uint64
}

func (b *BaseEntityImpl) TypeID() uint64 {
	return b.typeID
}

func (b *BaseEntityImpl) InjectTypeID(id uint64) {
	b.typeID = id
}
