package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0, 100),
	}
}

func (q *Queue) Enqueue(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}

	v := q.data[0]
	q.data = q.data[1:]

	return v, true
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		ptr1, ptr2 = l1, l2
		q1, q2     = NewQueue(), NewQueue()
	)

	for ptr1 != nil || ptr2 != nil {
		if ptr1 != nil && ptr2 != nil {
			q1.Enqueue(ptr1.Val)
			ptr1 = ptr1.Next

			q2.Enqueue(ptr2.Val)
			ptr2 = ptr2.Next

		} else if ptr1 != nil {
			q1.Enqueue(ptr1.Val)
			q2.Enqueue(0)

			ptr1 = ptr1.Next
		} else {
			q2.Enqueue(ptr2.Val)
			q1.Enqueue(0)

			ptr2 = ptr2.Next
		}
	}

	var mem int

	elem1, ok1 := q1.Dequeue()
	elem2, ok2 := q2.Dequeue()

	newList := &ListNode{
		Val: (elem1 + elem2) % 10,
	}
	tail := newList

	mem = (elem1 + elem2) / 10

	for {
		elem1, ok1 = q1.Dequeue()
		elem2, ok2 = q2.Dequeue()

		if !(ok1 && ok2) {
			break
		}

		tail.Next = &ListNode{}
		tail = tail.Next

		tail.Val = (elem1 + elem2 + mem) % 10
		mem = (elem1 + elem2 + mem) / 10
	}

	if mem > 0 {
		tail.Next = &ListNode{}
		tail = tail.Next

		tail.Val = mem
	}

	return newList
}
