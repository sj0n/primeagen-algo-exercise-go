package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestTwoCrystallBall(t *testing.T) {
	floors := [10000]bool{}
	t.Parallel()
	t.Run("Floor limit", func(t *testing.T) {
		index := int(math.Floor(rand.Float64() * 10000))

		for i := index; i < len(floors); i++ {
			floors[i] = true
		}

		floor_limit := TwoCrystalBall(floors[:])

		if floor_limit != index {
			t.Errorf("Expected %d", index)
			t.Errorf("Received %d", floor_limit)
		}
	})

	t.Run("No floor limit", func(t *testing.T) {
		floors = [10000]bool{}

		floor_limit := TwoCrystalBall(floors[:])

		if floor_limit != -1 {
			t.Errorf("Expected %d", -1)
			t.Errorf("Received %d", floor_limit)
		}
	})
}