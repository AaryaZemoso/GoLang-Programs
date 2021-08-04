package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p, q := Point{10, 20}, Point{50, 40}
	fmt.Println(p, q)

	// We can use . (dot) operator to access the properties/variables of the structs
	fmt.Println(p.X, p.Y)

	fmt.Println()

	// Can also use pointers to access the objects
	x := &p
	fmt.Println(x, *x)
	x.X = 10e3
	fmt.Println(x, *x)
	fmt.Println(p)

	fmt.Println()

	// Can also use struct Literals to initialize selected variables
	y := &Point{X: 10}
	fmt.Println(y, *y)
}
