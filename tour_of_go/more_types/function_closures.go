package main

import "fmt"

func add(initSum int) func(int) int {
	sum := initSum
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	alpha, beta := add(10), add(20)
	for i := 1; i < 10; i++ {
		fmt.Println(
			alpha(i),
			beta(i),
		)
	}

}
