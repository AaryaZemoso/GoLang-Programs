package main

import "fmt"

func main() {

	// While loop in Go is same as for but only conditional statement exists
	// Eg:-
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// NOTE: If for has no statements, its a forever loop/infinite loop
}
