package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Point struct {
	X, Y float64
}

// No explicit use of implement keyword like in java
// It assumes you are implementing the required interfaces implicitly
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{3, 4}

	var ab Abser
	ab = &p

	fmt.Println(ab.Abs())
}
