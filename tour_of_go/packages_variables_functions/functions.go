package main

import "fmt"

// params are given differently
// variable name then comes data type
// You can join consecutive params and give the common datatype
// eg: add(x int, y int) is same as add(x, y int)
func add(x, y int) int {
	return x + y
}

// You cannot have same functions name I guess.
// Polymorphism doesn't work???
// Though I'm changing the function signature to take 3 param its not working
// Had to change add(x, y, z int) -> add3(x, y, z int)
func add3(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add3(10, 20, 30))
}
