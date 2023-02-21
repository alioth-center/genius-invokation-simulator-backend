package definition

import "github.com/sunist-c/genius-invokation-simulator-backend/enum"

type Damage interface {
	ElementType() enum.ElementType
	DamageAmount() uint
}
