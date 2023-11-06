package stack

import "errors"

type Node[T any] struct {
	Value T
	Prev  *Node[T]
}

type Stack[T any] struct {
	Head   *Node[T]
	Length int
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{Value: value}
	s.Length++

	if s.IsEmpty() {
		s.Head = node
	} else {
		node.Prev = s.Head
		s.Head = node
	}
}

func (s *Stack[T]) Pop() (interface{}, error) {
	if s.Length != 0 {
		s.Length--
	}

	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	node := s.Head
	s.Head = node.Prev

	return node.Value, nil
}

func (s *Stack[T]) Peek() T {
	return s.Head.Value
}