package queue_on_2_pointers

import "errors"

type queue[T any] struct {
	data []T
	head int
	tail int
	size int
}

func New[T any](initialCap uint64) *queue[T] {
	if initialCap < 1 {
		initialCap = 1
	}
	return &queue[T]{
		data: make([]T, initialCap),
	}
}

func (q *queue[T]) PushBack(value T) {
	if q.size == len(q.data) {
		newData := make([]T, len(q.data)*2)

		for i := 0; i < q.size; i++ {
			newData[i] = q.data[(q.head+i)%len(q.data)]
		}

		q.data = newData
		q.head = 0
		q.tail = q.size
	}

	q.data[q.tail] = value
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *queue[T]) PopFront() (T, bool) {
	if q.size == 0 {
		var zero T

		return zero, false
	}

}

func (q *queue[T]) Empty() bool {
	return q.head == q.tail
}

func (q *queue[T]) Front() (T, bool) {

}
