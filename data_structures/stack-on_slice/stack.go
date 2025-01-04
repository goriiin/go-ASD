package stack_on_slice

import "errors"

type stack[T any] struct {
	data []T
}

func New[T any]() *stack[T] {
	return &stack[T]{data: make([]T, 0)}
}

func (s *stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		var zero T

		return zero, errors.New("stack is empty")
	}

	ans := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return ans, nil
}

func (s *stack[T]) Top() (T, error) {
	if len(s.data) == 0 {
		var zero T

		return zero, errors.New("stack is empty")
	}

	ans := s.data[len(s.data)-1]

	return ans, nil
}

func (s *stack[T]) Empty() bool {
	return len(s.data) == 0
}
