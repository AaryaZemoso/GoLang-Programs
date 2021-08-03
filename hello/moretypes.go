package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func MoreTypes() {

	// Pointers
	fmt.Println("\n\nMore Types:")
	fmt.Println("\nPointer")

	var p *int
	i := 10
	p = &i

	fmt.Println(p, *p)

	*p = 20
	fmt.Println(p, *p)

	// Structs
	fmt.Println("Structs")
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 10
	fmt.Println(v, v.X)

	// Arrays
	fmt.Println("Arrays")
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i * 2
	}
	fmt.Println(a)

	// Slices
	fmt.Println("Slices")
	fmt.Println(a[3:6])

	fmt.Println("Length:", len(a[4:]))
	fmt.Println("Capacity:", cap(a[4:]))

	// Dynamic Slice/Array
	fmt.Println("Dynamic Array")
	n := 5
	arr := make([]int, n)
	fmt.Println(arr)
	arr = append(arr, 100)
	fmt.Println(arr)

	// Range
	fmt.Println("Range")
	var pow = []int{2, 4, 8, 16}
	for i, v := range pow {
		fmt.Println(i, v)
	}

	// Maps
	fmt.Println("Maps")
	m := make(map[string]int)
	m["hi"] = 10
	m["aarya"] = 20
	fmt.Println(m["hi"], m)

	delete(m, "aarya")
	ele, ok := m["aarya"]
	fmt.Println(ele, ok)

}
