package BFS

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type QueueNode[T any] struct {
	Node *TreeNode[T]
	Next *QueueNode[T]
}

type Queue[T any] struct {
	Head   *QueueNode[T]
	Tail   *QueueNode[T]
	Length int
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue[T]) Enqueue(treeNode *TreeNode[T]) {
	node := &QueueNode[T]{Node: treeNode}
	q.Length++

	if q.IsEmpty() {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue[T]) Dequeue() *QueueNode[T] {
	if q.IsEmpty() {
		return nil
	}

	node := q.Head
	q.Head = q.Head.Next

	return node
}

func BFS(root *TreeNode[int], needle int) bool {
	node := &QueueNode[int]{Node: root}
	q := Queue[int]{Head: node}

	for !q.IsEmpty() {
		current := q.Dequeue()

		if current.Node.Value == needle {
			return true
		}

		if current.Node.Left != nil {
			q.Enqueue(current.Node.Left)
		}

		if current.Node.Right != nil {
			q.Enqueue(current.Node.Right)
		}
	}
	return false
}
