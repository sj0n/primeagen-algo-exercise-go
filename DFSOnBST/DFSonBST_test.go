package dfs

import "testing"

func TestDFS(t *testing.T) {
	root := &BinaryNode{Value: 43,
		Left: &BinaryNode{
			Value: 22,
			Left: &BinaryNode{
				Value: 18,
				Left: &BinaryNode{
					Value: 14,
				},
				Right: &BinaryNode{
					Value: 20,
				},
			},
			Right: &BinaryNode{
				Value: 30,
				Left: &BinaryNode{
					Value: 27,
				},
				Right: &BinaryNode{
					Value: 41,
				},
			},
		},
		Right: &BinaryNode{
			Value: 57,
			Left: &BinaryNode{
				Value: 49,
			},
			Right: &BinaryNode{
				Value: 61,
			},
		},
	}

	testCases := []struct {
		a        *BinaryNode
		b        int
		expected bool
	}{
		{root, 61, true},
		{root, 1, false},
	}

	for _, test := range testCases {
		if result := dfs(test.a, test.b); result != test.expected {
			t.Errorf("Expected %v, Received %v", test.expected, result)
		}
	}
}
