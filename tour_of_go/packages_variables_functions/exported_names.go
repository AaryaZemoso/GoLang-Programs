package main

import (
	"fmt"
	"math"
)

func main() {
	// Names are exported only if the first letter is Capital
	// math.Pi is legal
	// math.pi is not legal
	fmt.Println(math.Pi)
}
