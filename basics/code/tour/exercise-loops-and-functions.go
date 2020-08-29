package main

import (
	"fmt"
	"math"
)

const delta = 1.0e-10

func Sqrt(x float64) float64 {
	/* function to be implemented */
	z := 1.0 // guess
	// z := (x / 2.0) // guess 2 -> works
	// z := x // guess 3 -> works
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Sqrt: %v", z)
		fmt.Printf("\tDifference: %v\n", z-math.Sqrt(x)) // Difference from actual Sqrt
		if diff := z - math.Sqrt(x); diff < delta {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(5))
}
