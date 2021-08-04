package main

import "fmt"

// Return variable can be defined before itself and
// giving an empty return statement will return the variables defined at the function signature
// This is also called as "Naked" return
func add10(s int) (y int) {
	y = s + 10
	return
}

func main() {
	fmt.Println(add10(10))
}
