package main

import "math"

func TwoCrystalBall(floors []bool) int {
	jumpAmount := int(math.Floor((math.Sqrt(float64(len(floors))))))
	i := jumpAmount
	ballBroken := false

	for ; jumpAmount < len(floors); i += jumpAmount {
		if (i >= len(floors)) {
			break
		}
		
		if floors[i] {
			ballBroken = true
			break
		}
	}

	if ballBroken {
		i -= jumpAmount
		for ; i < len(floors); i++ {
			if (floors[i]) {
				return i
			}
		}
	}

	return -1
}