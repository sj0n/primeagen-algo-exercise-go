package btinorder

import (
	"reflect"
	"testing"
)

func TestBTInOrder(t *testing.T) {
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

	treeSlice := InOrderSearch(tree)
	result := reflect.DeepEqual(treeSlice, []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100})

	if !result {
		t.Errorf("Expected %v", []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100})
		t.Errorf("Received %v", treeSlice)
	}
}
