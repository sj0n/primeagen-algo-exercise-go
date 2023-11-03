package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestFloorLimit(t *testing.T) {
	index := int(math.Floor(rand.Float64() * 10000))
	floors := [10000]bool{}
	
	for i := index; i < len(floors); i++ {
		floors[i] = true
	}

	floor_limit := TwoCrystalBall(floors[:]);

	if floor_limit != index {
		t.Errorf("Expected %d", index)
		t.Errorf("Received %d", floor_limit)
	}
}

func TestNoFloorLimit(t *testing.T) {
	floors := [10000]bool{}

	floor_limit := TwoCrystalBall(floors[:])

	if floor_limit != -1 {
		t.Errorf("Expected %d", -1)
		t.Errorf("Received %d", floor_limit)
	}
}