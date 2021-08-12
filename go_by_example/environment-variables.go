package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// for _, e := range os.Environ() {
	// 	fmt.Println(e)
	// }

	path_string := os.Getenv("PATH")

	for _, e := range strings.SplitN(path_string, ":", -1) {
		fmt.Println(e)
	}
}
