package main

import (
	"fmt"
	"math"
)

const delta = 1e-10

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// fmt.Printf("type: %T\n", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x // guess

	for {
		n := z - (z*z-x)/(2*z)
		fmt.Printf("Sqrt: %v", z)
		fmt.Printf("\tDifference: %v\n", math.Abs(n-z))
		if diff := math.Abs(n - z); diff < delta {
			break
		}
		z = n
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
