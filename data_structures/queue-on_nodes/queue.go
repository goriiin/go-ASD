package queue_on_nodes

type node[T any] struct {
	data T
	next *node[T]
}

func NewNode[T any](data T, next *node[T]) *node[T] {
	return &node[T]{
		data: data,
		next: next,
	}
}

func (n *node[T]) Set(data T, next *node[T]) {
	n.data = data
	n.next = next
}

func (n *node[T]) SetNext(next *node[T]) {
	n.next = next
}

func (n *node[T]) Value() T {
	return n.data
}

func (n *node[T]) Next() *node[T] {
	return n.next
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *queue[T] {
	return &queue[T]{
		head: nil,
		tail: nil,
	}
}

func (q *queue[T]) PushBack(data T) {
	newNode := NewNode(data, nil)

	if q.Empty() {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.SetNext(newNode)
	q.tail = newNode
}

func (q *queue[T]) PopFront() (T, bool) {
	if q.Empty() {
		var zero T

		return zero, false
	}

	ans := q.head.Value()
	q.head = q.head.Next()

	if q.head == nil {
		q.tail = nil
	}

	return ans, true
}

func (q *queue[T]) Front() (T, bool) {
	if q.Empty() {
		var zero T

		return zero, false
	}

	return q.head.Value(), true
}

func (q *queue[T]) Empty() bool {
	return q.head == nil
}
