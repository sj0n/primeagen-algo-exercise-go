package main

import "testing"

func TestNumberFound(t *testing.T) {
	index := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)

	if index != 6 {
		t.Errorf("Got %d, want %d", index, 6)
	}
}

func TestNumberNotFound(t *testing.T) {
	index := BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 420)

	if index != -1 {
		t.Errorf("Got %d, want %d", index, -1);
	}
}