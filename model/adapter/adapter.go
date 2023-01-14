package adapter

type Adapter[Origin any, Dest any] interface {
	Convert(source Origin) (success bool, result Dest)
}
