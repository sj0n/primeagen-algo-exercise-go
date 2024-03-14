package btpostorder

import (
	"reflect"
	"testing"
)

func TestBTPostOrder(t *testing.T) {
	tree := &BinaryNode[int]{
		value: 20,
		left: &BinaryNode[int]{
			value: 10,
			left: &BinaryNode[int]{
				value: 5,
				right: &BinaryNode[int]{
					value: 7},
			},
			right: &BinaryNode[int]{
				value: 15},
		},
		right: &BinaryNode[int]{
			value: 50,
			left: &BinaryNode[int]{
				value: 30,
				left: &BinaryNode[int]{
					value: 29,
				},
				right: &BinaryNode[int]{
					value: 45,
				},
			},
			right: &BinaryNode[int]{
				value: 100,
			},
		},
	}

	treeSlice := PostOrderTraverse(tree)
	result := reflect.DeepEqual(treeSlice, []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20})

	if !result {
		t.Errorf("Expected %v", []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20})
		t.Errorf("Received %v", treeSlice)
	}
}
