package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var (
		st  = NewStack()
		ptr = root
		ok  bool

		res = make([]int, 0, 100)

		prev *TreeNode
	)

	if ptr == nil {
		return res
	}

	st.Push(ptr)

	for len(st.data) > 0 {
		ptr, ok = st.Peek()
		if !ok {
			break
		}

		if prev == nil || prev.Right == ptr || prev.Left == ptr {
			if ptr.Left != nil {
				st.Push(ptr.Left)
			} else if ptr.Right != nil {
				st.Push(ptr.Right)
			} else {
				st.Pop()

				res = append(res, ptr.Val)
			}
		} else if ptr.Left == prev {
			if ptr.Right != nil {
				st.Push(ptr.Right)
			} else {
				st.Pop()
				res = append(res, ptr.Val)
			}
		} else {
			res = append(res, ptr.Val)

			st.Pop()
		}

		prev = ptr
	}

	return res
}

type Stack struct {
	data []*TreeNode
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*TreeNode, 0, 100),
	}
}

func (s *Stack) Push(val *TreeNode) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (*TreeNode, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return val, true
}

func (s *Stack) Peek() (*TreeNode, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	return s.data[len(s.data)-1], true
}
