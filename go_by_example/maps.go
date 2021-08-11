package main

import "fmt"

func main() {

	// map[key-type]value-type
	m := make(map[string]int)

	// Setting values
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Getting values
	v1 := m["k1"]

	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// Using second return value to check if key exists
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declaration and initialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
