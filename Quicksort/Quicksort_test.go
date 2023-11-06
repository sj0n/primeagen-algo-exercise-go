package quicksort

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{9, 3, 7, 4, 69, 420, 42}
	QuickSort(array, 0, len(array) - 1)

	result := sort.SliceIsSorted(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	if !result {
		t.Errorf("Expected %v", []int{3, 4, 7, 9, 42, 69, 420})
		t.Errorf("Received %v", array)
	} else {
		fmt.Printf("Expected %v\n", []int{3, 4, 7, 9, 42, 69, 420})
		fmt.Printf("Received %v\n", array)
	}
}
