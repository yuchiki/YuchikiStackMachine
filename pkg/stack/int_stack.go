package stack

type Uint64Stack interface {
	Push(n uint64)

	Pop() (uint64, error)
}

func NewUint64Stack() Uint64Stack {
	return &uint64Stack{}
}
