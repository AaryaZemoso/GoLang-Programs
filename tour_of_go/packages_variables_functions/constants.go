package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println(Pi)

	// Following is illegal since Pi is constant and cannot be changed
	// Pi = 2.534
	// fmt.Println(Pi)

	const BIG = 1<<63 - 1
	fmt.Println(BIG)

}
