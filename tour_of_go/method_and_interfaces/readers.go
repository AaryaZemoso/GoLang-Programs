package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Reader")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err, b)
		fmt.Printf("%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
