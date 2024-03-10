package BFS

import "testing"

func TestBFS(t *testing.T) {
	root := &TreeNode[int]{Value: 3}
	root.Left = &TreeNode[int]{Value: 4}
	root.Right = &TreeNode[int]{Value: 5}
	root.Left.Left = &TreeNode[int]{Value: 1}
	root.Left.Right = &TreeNode[int]{Value: 2}

	t.Run("Test for true", func(t *testing.T) {
		result := BFS(root, 3)

		if !result {
			t.Errorf("Expected true, Received %v", result)
		}
	})

	t.Run("Test for false", func(t *testing.T) {
		result := BFS(root, 6)

		if result {
			t.Errorf("Expected false, Received %v", result)
		}
	})
}
