package main

import "fmt"

// Pointers are similar to C
// But there is no pointer arithmetic

func main() {
	i, j := 1, 2

	p := &i

	fmt.Println(p, *p)

	p = &j

	fmt.Println(p, *p)

	// Any changes can also be done using the pointer
	// Also called indirecting or dereferencing

	*p = 20
	fmt.Println(p, *p)
}
