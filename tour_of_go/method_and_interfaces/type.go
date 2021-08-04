package main

import "fmt"

func main() {

	var i interface{}

	i = 42
	fmt.Println(i)

	i = "Aarya"
	fmt.Println(i)

	// Type Assertions
	t, ok := i.(string)
	fmt.Println(t, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Type Switches
	switch i.(type) {
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float")
	}
}
