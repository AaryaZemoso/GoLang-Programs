package main

import "fmt"

func main() {
	// Map matching with id to name
	emp := make(map[int]string)

	// Inserting
	emp[1] = "Aarya"
	emp[2] = "Devarla"

	// Deleting
	delete(emp, 1)

	// Test if a key is present or not
	ele, ok := emp[1]
	if ok {
		fmt.Println(ele)
	} else {
		fmt.Println("Key doesnt exist")
	}

	fmt.Println(emp)
}
