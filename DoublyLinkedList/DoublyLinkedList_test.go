package doublylinkedlist

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	list := DoublyLinkedList[int]{}
	t.Run("Append()", func(t *testing.T) {
		list.Append(1)
		list.Append(2)
		list.Append(3)
		t.Run("Length is 3", func(t *testing.T) {
			if list.Length != 3 {
				t.Errorf("Expected %d", 3)
				t.Errorf("Received %d", list.Length)
			}
		})

		t.Run("Head value is 1", func(t *testing.T) {
			if list.Head.Value != 1 {
				t.Errorf("Expected %d", 1)
				t.Errorf("Received %d", list.Head.Value)
			}
		})

		t.Run("Tail value is 3", func(t *testing.T) {
			if list.Tail.Value != 3 {
				t.Errorf("Expected %d", 3)
				t.Errorf("Received %d", list.Tail.Value)
			}
		})
	})

	t.Run("GetAt() element at index 1", func(t *testing.T) {
		if node, _ := list.GetAt(1); node.Value != 2 {
			t.Errorf("Expected %d", 2)
			t.Errorf("Received %d", node.Value)
		}
	})

	t.Run("RemoveAt()", func(t *testing.T) {
		t.Run("Remove 0 index element returns the element", func(t *testing.T) {
			node, _ := list.RemoveAt(0)

			if node.Value != 1 {
				t.Errorf("Expected %d", 1)
				t.Errorf("Received %d", node.Value)
			}
		})

		t.Run("Length is 2", func(t *testing.T) {
			if list.Length != 2 {
				t.Errorf("Expected %d", 2)
				t.Errorf("Received %d", list.Length)
			}
		})
	})

	list.RemoveAt(0)
	list.RemoveAt(0)

	t.Run("Length is 0", func(t *testing.T) {
		if list.Length != 0 {
			t.Errorf("Expected %d", 0)
			t.Errorf("Received %d", list.Length)
		}
	})

	t.Run("Prepend()", func(t *testing.T) {
		list.Prepend(1)
		list.Prepend(2)
		list.Prepend(3)

		t.Run("Length is 3", func(t *testing.T) {
			if list.Length != 3 {
				t.Errorf("Expected %d", 3)
				t.Errorf("Received %d", list.Length)
			}
		})

		t.Run("Value of head is 3", func(t *testing.T) {
			if list.Head.Value != 3 {
				t.Errorf("Expected %d", 3)
				t.Errorf("Received %d", list.Head.Value)
			}
		})

		t.Run("Value of tail is 1", func(t *testing.T) {
			if list.Tail.Value != 1 {
				t.Errorf("Expected %d", 1)
				t.Errorf("Received %d", list.Tail.Value)
			}
		})
	})
}
