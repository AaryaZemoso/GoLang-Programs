package main

import "fmt"

// Self explanatory
// Similar to Python
func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := 10, 20
	fmt.Println("a =", a, "b =", b)
	a, b = swap(a, b)
	fmt.Println("a =", a, "b =", b)
}
