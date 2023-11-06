package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := Queue[int]{}
	t.Run("Enqueue(): Length of the Queue is 3", func(t *testing.T) {
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		if queue.Length != 3 {
			t.Errorf("Expected %d, Received %d", 3, queue.Length)
		}
	})

	t.Run("Dequeue()", func(t *testing.T) {
		t.Parallel()
		queue.Dequeue()
		t.Run("Queue length reduce by 1", func(t *testing.T) {
			if queue.Length != 2 {
				t.Errorf("Expected %d, Received %d", 2, queue.Length)
			}
		})

		t.Run("Peek() Value of head updates", func(t *testing.T) {
			if queue.Peek() != 2 {
				t.Errorf("Expected %d, Received %d", 2, queue.Peek())
			}
		})

		t.Run("Dequeue() on empty queue returns error", func(t *testing.T) {
			queue.Dequeue()
			queue.Dequeue()
			_, error := queue.Dequeue()

			if error == nil {
				t.Errorf("Expected %v, Received %v", "Queue is empty", error)
			}
		})
	})
}