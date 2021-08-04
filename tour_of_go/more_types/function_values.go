package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(40, 50)
}

func main() {
	// Passing Multiplication function as a parameter
	fmt.Println(compute(func(x, y int) int {
		return x * y
	}))

	// Passing Addition function as a parameter
	fmt.Println(compute(func(x, y int) int {
		return x + y
	}))
}
