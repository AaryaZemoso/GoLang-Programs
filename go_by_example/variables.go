package main

import "fmt"

func main() {

	// Variable Initialization
	var a = "initial"
	fmt.Println(a)

	// Multiple Variable Initialization
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Values are zeroed if not assigned
	var e int
	fmt.Println(e)

	// Shorthand
	f := "apple"
	fmt.Println(f)

}
