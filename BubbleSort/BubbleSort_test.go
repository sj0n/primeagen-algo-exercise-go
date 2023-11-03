package bubblesort

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{9, 3, 7, 4, 69, 420, 42}
	BubbleSort(array)

	result := sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	if !result {
		t.Errorf("Expected %v, Received %v", []int{3, 4, 7, 9, 42, 69, 420} , array)
	} else {
		fmt.Printf("Expected %v, Received %v\n", []int{3, 4, 7, 9, 42, 69, 420} , array)
	}
}