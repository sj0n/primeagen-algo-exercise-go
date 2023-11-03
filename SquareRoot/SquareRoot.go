package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x/2
	prev := z
	prevDiff := 0.0
	currentDiff := 0.0
	
	for i := 0; i < 15; i++ {
		z -= (z*z - x) / (2 * z)
		currentDiff = z - prev
		
		if i != 0 && currentDiff >= prevDiff {
			return z
		}
		
		fmt.Println(z)
		prevDiff = currentDiff
	}
	return z
}

func main() {
	fmt.Println("\nLoop sqrt:", Sqrt(250))
	fmt.Println("Math sqrt:", math.Sqrt(250))
}
