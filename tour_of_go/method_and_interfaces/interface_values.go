package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

// Abstraction ??????
// Able to access the method using the reference of the object

func main() {

	t := T{"Aarya"}
	f := F(math.Pi)

	var i I

	i = &t
	i.M()

	i = &f
	i.M()

}
