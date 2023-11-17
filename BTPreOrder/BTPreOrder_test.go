package btpreorder

import (
	"reflect"
	"testing"
)

func TestBTPreOrder(t *testing.T) {
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

	treeSlice := PreOrderSearch(tree)
	result := reflect.DeepEqual(treeSlice, []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100})

	if !result {
		t.Errorf("Expected %v", []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100})
		t.Errorf("Received %v", treeSlice)
	}

}
