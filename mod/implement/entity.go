package implement

type BaseEntityImpl struct {
	typeID uint
}

func (b *BaseEntityImpl) TypeID() uint {
	return b.typeID
}

func (b *BaseEntityImpl) InjectTypeID(id uint) {
	b.typeID = id
}
