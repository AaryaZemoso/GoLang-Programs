package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func main() {
	p := Point{10, 20}

	fmt.Println(p)
	p.Scale(0.5)
	fmt.Println(p)

}
