package main

import "fmt"

// Outside function
// It must follow a strict syntax rule
var windows, mac, linux string = "Microsoft", "Apple", "Linus"

// This is illegal
// c := 100

func main() {

	// Inside function
	// You can use short variable declarations
	c := 10
	fmt.Println(c)

	var java int = 50
	fmt.Println(java)
}
