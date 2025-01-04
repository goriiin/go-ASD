package interfaces

type Queue[T any] interface {
	PushBack(value T)
	PopFront() (T, bool)
	Front() (T, bool)
	Empty() bool
}
