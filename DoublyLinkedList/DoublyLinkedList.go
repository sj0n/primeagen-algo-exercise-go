package doublylinkedlist

import "errors"

type Node[T any] struct {
	Prev  *Node[T]
	Next  *Node[T]
	Value T
}

type DoublyLinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.Head == nil
}

func (dll *DoublyLinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value}
	dll.Length++

	if dll.IsEmpty() {
		dll.Head = node
		dll.Tail = node
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DoublyLinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}
	dll.Length++

	if dll.IsEmpty() {
		dll.Head = node
		dll.Tail = node
		return
	}

	node.Prev = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node
}

func (dll *DoublyLinkedList[T]) GetAt(index int) (*Node[T], error) {
	currentNode := dll.Head

	if currentNode == nil {
		return nil, errors.New("head is empty")
	}

	for i := 0; currentNode != nil && i < index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode, nil
}

func (dll *DoublyLinkedList[T]) InsertAt(value T, index int) error {
	if index < 0 || index > dll.Length {
		return errors.New("index is out of bound")
	}

	if index == dll.Length {
		dll.Append(value)
	} else if index == 0 {
		dll.Prepend(value)
	} else {
		currentNode, error := dll.GetAt(index)

		if error != nil {
			return error
		}

		dll.Length++
		node := &Node[T]{Value: value}

		node.Next = currentNode
		node.Prev = currentNode.Prev
		currentNode.Prev = node

		if node.Prev != nil {
			node.Prev.Next = node
		}
	}
	return nil
}

func (dll *DoublyLinkedList[T]) RemoveAt(index int) (*Node[T], error) {
	if index < 0 || index > dll.Length {
		return nil, errors.New("index is out of bound")
	}

	node, error := dll.GetAt(index)

	if error != nil {
		return nil, error
	}

	dll.Length--

	if dll.Length == 0 {
		node := dll.Head
		dll.Head = nil
		dll.Tail = nil

		return node, nil
	}

	if node.Prev != nil && node.Next != nil {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	if node == dll.Head {
		dll.Head = node.Next
	}

	if node == dll.Tail {
		dll.Tail = node.Prev
	}

	return node, nil
}
