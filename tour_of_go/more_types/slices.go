package main

import (
	"fmt"
	"strings"
)

func main() {

	// Slices are a dynamic flexible view of an array

	// Array:-
	a := [5]int{1, 2, 3, 4, 5}

	// Slice:-
	b := a[1:3]
	fmt.Println(a, b)

	// They are references to the arrays
	// Any changes to the slice is reflected on the original Array itself.
	b[0] = 10
	fmt.Println(a, b)

	fmt.Print("a = ")
	fmt.Println(len(a), cap(a))

	fmt.Print("b = ", b)
	fmt.Println(len(b), cap(b))

	b = b[:4]
	fmt.Print("extened b = ", b)
	fmt.Println(len(b), cap(b))

	fmt.Println()

	// Using make()
	c := make([]int, 4)
	fmt.Println(c)

	c[2] = 3
	c = c[1:3]
	fmt.Println(c)

	fmt.Println()

	// Slice of slices
	d := [][]string{
		[]string{"10", "20"},
		[]string{"20", "30"},
	}
	
	for i := range d {
		fmt.Println(strings.Join(d[i], " "))
	}

	fmt.Println()

	// Appending to slice
	d[0] = append(d[0], "30")
	fmt.Println(d)

}
