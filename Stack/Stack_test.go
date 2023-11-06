package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Stack[int]{}
	t.Run("Length of stack is 3", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.Length != 3 {
			t.Errorf("Expected %d", 3)
			t.Errorf("Received %d", stack.Length)
		}
	})

	t.Run("Value on top of the stack is the last in", func(t *testing.T) {
		if stack.Peek() != 3 {
			t.Errorf("Expected %d", 3)
			t.Errorf("Received %d", stack.Peek())
		}
	})

	t.Run("Pop()", func(t *testing.T) {
		value, _ := stack.Pop()
		t.Parallel()
		t.Run("Length is 2", func(t *testing.T) {

			if stack.Length != 2 {
				t.Errorf("Expected %d", 2)
				t.Errorf("Received %d", stack.Length)
			}
		})

		t.Run("Value popped should be last in", func(t *testing.T) {
			if value != 3 {
				t.Errorf("Expected %d", 3)
				t.Errorf("Received %d", value)
			}
		})
	})
}