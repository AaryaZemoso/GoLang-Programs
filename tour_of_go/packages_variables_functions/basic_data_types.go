package main

import "fmt"

/*
	Basic types provided are:-

	bool

	string

	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte - alias to uint8
	rune - alias to int32

	float32 float 64

	complex64 complex128

*/

// If no value is initialized the variables are given their repective zero value
var (
	vbool    bool    = false
	vstring  string  = "some string"
	vint     int     = 10
	vfloat   float64 = 12.4563
	vbyte    byte
	vcomplex complex64 = 5 + 2i
)

func main() {
	fmt.Println(vbool)
	fmt.Println(vbyte)
	fmt.Println(vcomplex)
	fmt.Println(vfloat)
	fmt.Println(vint)
	fmt.Println(vstring)
}
