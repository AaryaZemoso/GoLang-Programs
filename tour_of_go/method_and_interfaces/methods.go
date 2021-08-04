package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{30, 40}
	fmt.Println(p.Abs())
	fmt.Println(Abs(p))
}
