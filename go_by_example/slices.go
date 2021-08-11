package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Array like operations
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Additional functions
	// Append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	// Copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slicing
	l := s[2:5]
	fmt.Println("sli:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// Declaration and Initialization - Shorthand
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// 2D Array
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
