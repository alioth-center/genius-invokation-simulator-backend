package implement

type BaseEntityImplement struct {
	typeID uint
}

func (b *BaseEntityImplement) TypeID() uint {
	return b.typeID
}

func (b *BaseEntityImplement) InjectTypeID(id uint) {
	b.typeID = id
}
