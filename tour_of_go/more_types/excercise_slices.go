package main

import "golang.org/x/tour/pic"

// Unable to run in the local system
// No giving image as output
func Pic(dx, dy int) [][]uint8 {
	var arr = make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = uint8(i * j)
		}
	}

	return arr
}

func main() {
	pic.Show(Pic)
}
