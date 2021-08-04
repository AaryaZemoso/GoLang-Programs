package main

import "fmt"

func main() {

	// arrays 2 ways of declaring
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println(arr)

	a := [5]int{1, 3, 5, 7, 9}
	fmt.Println(a)
}
