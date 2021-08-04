package main

import "fmt"

func main() {
	i := 10

	// No need of paranthesis
	// The curly brackets should follow immediately after the condition and cannot be placed in next line.
	// Because Go compiler assumes it is end of statement and throws an error.

	// Similar with else, it should follow the ending of the if curly brackets and not in next line.
	if i < 20 {
		fmt.Println("Less than 20")
	} else {
		fmt.Println("More than 20")
	}

	// If with Short Statement
	// if can have an initialization statement or short statement which executes before condition is checked
	// Variables declared in this short statement are accessible in the scope of if and else
	if j := 20; i < j {
		fmt.Println(i, "<", j)
	} else {
		fmt.Println(i, ">", j)
	}
}
