package main

import "fmt"

// No implicit type conversion
// Need to use explicit type conversion

func main() {

	i := 6
	var x float64 = 34.20
	var f float64 = float64(i) * x

	fmt.Println(f)

	var ui uint64 = uint64(f)
	fmt.Println(ui)
}
