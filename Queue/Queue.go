package queue

import "errors"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{Value: value}
	q.Length++

	if q.IsEmpty() {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue[T]) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	
	if !q.IsEmpty() {
		q.Length--
	}

	value := q.Head.Value
	q.Head = q.Head.Next

	return value, nil
}

func (q *Queue[T]) Peek() T {
	return q.Head.Value
}