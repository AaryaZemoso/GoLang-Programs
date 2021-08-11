package main

import "fmt"

func main() {

	// While loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// For loop
	for j := 7; j <= 9; j++{
		fmt.Println(j)
	}

	// Infinite loop unless it sees break
	for{
		fmt.Println("loop")
		break
	}

	// Using continue
	for n := 0; n <= 5; n++{
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	}

}
