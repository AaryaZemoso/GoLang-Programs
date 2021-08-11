package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, nums := range nums {
		total += nums
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
