package main

import "testing"

var numArray = []int{3, 17, 21, 1, 45, 13}

func TestNumberFound(t *testing.T) {
	var index = linearSearch(numArray, 21)

	if index == 2 {
		t.Errorf("Got %d, want %d", index, 2);
	}
}

func TestNumberNotFound(t *testing.T) {
	var index = linearSearch(numArray, 55)

	if index != -1 {
		t.Errorf("Got %d, want %d", index, -1)
	}
}