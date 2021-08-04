package main

import "fmt"

func AddFunc(x, y int) {
	// This statement executes before this function(AddFunc) is returned
	// Normally used to write the code for closing of files (and sessions?? maybe)
	defer fmt.Println("Addition Successful")
	fmt.Println(x + y)
}

func stackDefer() {
	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done with execution")
	fmt.Println("Defer statements will now begin executing")
	fmt.Println("They are stored in stack form")
}

func main() {
	AddFunc(10, 54)
	stackDefer()
}
