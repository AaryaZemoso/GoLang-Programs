package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// returns index, value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// returns key, value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// returns key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// returns index, character
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
