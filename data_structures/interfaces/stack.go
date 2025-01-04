package interfaces

type Stack[T any] interface {
	Push(v T)
	Pop() (T, error)
	Top() (T, error)
	Empty() bool
}
