package model

// BaseEntity 实体的基本接口，包含被框架托管的必须方法
type BaseEntity interface {
	TypeID() uint64
	InjectTypeID(id uint64)
}
