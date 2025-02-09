package main

type Stack struct {
	data []*TreeNode
}

func New() *Stack {
	return &Stack{data: make([]*TreeNode, 0, 200)}
}

func (s *Stack) Push(v *TreeNode) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (*TreeNode, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type Queue struct {
	in  *Stack
	out *Stack
}

func NewQueue() *Queue {
	return &Queue{in: New(), out: New()}
}

func (q *Queue) Enqueue(v *TreeNode) {
	q.in.Push(v)
}

func (q *Queue) Dequeue() (*TreeNode, bool) {
	if q.in.Empty() && q.out.Empty() {
		return nil, false
	} else if q.out.Empty() {
		for !q.in.Empty() {
			v, ok := q.in.Pop()
			if ok {
				q.out.Push(v)
			}
		}
	}

	return q.out.Pop()
}

func (q *Queue) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := NewQueue()
	q.Enqueue(root.Left)
	q.Enqueue(root.Right)

	for !q.Empty() {
		l, _ := q.Dequeue()
		r, _ := q.Dequeue()
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		q.Enqueue(l.Left)
		q.Enqueue(r.Right)
		q.Enqueue(l.Right)
		q.Enqueue(r.Left)
	}

	return true
}
