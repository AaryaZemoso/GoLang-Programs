package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	s = "md5 this string"

	h = md5.New()

	h.Write([]byte(s))

	bs = h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
