package queue_on_2_stacks

import (
	"github.com/goriiin/go-algo/data_structures/interfaces"
)

type queue[T any] struct {
	in  interfaces.Stack[T]
	out interfaces.Stack[T]
}

func New[T any](in, out interfaces.Stack[T]) *queue[T] {
	return &queue[T]{
		in:  in,
		out: out,
	}
}

func (q *queue[T]) PushBack(value T) {
	q.in.Push(value)
}

func (q *queue[T]) PopFront() (T, bool) {
	if q.Empty() {
		var zero T

		return zero, false
	}

	for !q.in.Empty() {
		item, _ := q.in.Pop()

		q.out.Push(item)
	}

	ans, err := q.out.Pop()
	if err != nil {
		var zero T

		return zero, false
	}

	return ans, true
}

func (q *queue[T]) Front() (T, bool) {
	if q.Empty() {
		var zero T

		return zero, false
	}

	for !q.in.Empty() {
		item, _ := q.in.Pop()

		q.out.Push(item)
	}

	ans, err := q.out.Top()
	if err != nil {
		var zero T

		return zero, false
	}

	return ans, true
}

func (q *queue[T]) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}
