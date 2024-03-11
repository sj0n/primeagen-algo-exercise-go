package comparebt

import "testing"

func TestCompareBT(t *testing.T) {
	root1 := &BinaryNode{Value: 1}
	root1.Left = &BinaryNode{Value: 2}
	root1.Right = &BinaryNode{Value: 3}

	root2 := &BinaryNode{Value: 1}
	root2.Left = &BinaryNode{Value: 2}
	root2.Right = &BinaryNode{Value: 3}

	root3 := &BinaryNode{Value: 2}
	root3.Left = &BinaryNode{Value: 1}

	root4 := &BinaryNode{Value: 2}
	root4.Right = &BinaryNode{Value: 1}

	testCases := []struct {
		a, b   *BinaryNode
		result bool
	}{
		{root1, root2, true},
		{root3, root4, false},
	}

	for _, tc := range testCases {
		result := compare(tc.a, tc.b)
		if result != tc.result {
			t.Errorf("Expected %v, Received %v", tc.result, result)
		}
	}
}
