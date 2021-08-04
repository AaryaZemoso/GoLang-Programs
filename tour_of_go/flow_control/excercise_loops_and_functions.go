package main

import (
	"fmt"
	"math"
)

// Implement this method based on the problem statement
func Sqrt(x float64) (z float64) {
	z = x / 2

	d := 10e-6

	for {
		z -= (z*z - x) / (x * 2)
		if math.Abs(z-math.Sqrt(x)) < d {
			break
		}
	}

	return
}

func main() {
	fmt.Println(Sqrt(2))
}
