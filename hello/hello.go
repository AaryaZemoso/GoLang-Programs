package main

import (
	"Zemoso/Bootcamp/GoLang/src/hello/hellostrings"
	"fmt"
	"math"
)

// Variables declared outside functions must follow this syntax
var p int = 100

// Syntax is a bit different from C Language
// Return type comes after giving the function signature

// Grouping of common datatypes and giving the datatype at end is allowed
// add(x int, y int) is same as add(x, y int)

// First variable -> then datatype
func add(x, y int) int {
	return x + y
}

// Multiple results can be returned
func swap(a, b int) (int, int) {
	return b, a
}

// Naked return
func fibNext(f, s int) (t, u int) {

	t = f + s
	u = s + t

	return
}

func main() {

	fmt.Println("hello from main")
	hellostrings.SayHelloFromPackage()

	// Names are exported if it begins with a Capital Letter
	fmt.Println("Exported names test")
	fmt.Println(math.Pi)

	// Functions
	fmt.Println("\nUsing of functions")
	fmt.Println(add(10, 15))

	// Multiple return values
	fmt.Println("\nMultiple return values")
	a, b := swap(10, 20)
	fmt.Println(a, b)

	// Naked return
	fmt.Println("Use of naked return")
	third, fourth := fibNext(2, 3)
	fmt.Println(third, fourth)

	// Declaring variables inside function
	// It does not need explicit use of keyword var
	// var x int = 10; ---> same as -----> x := 10
	fmt.Println("Variables in functions")
	var x int = 10
	fmt.Println(x)

	y := 20
	fmt.Println(y)

	fmt.Println("Variable outside function")
	fmt.Println(p)

	// Type conversions are explicit
	// NOTE:- The operator := is used for delcaration and initialization
	// 			where as = is used for assigning value
	fmt.Println("Type Conversions")
	var o = 42.5
	var q uint32 = uint32(o)
	fmt.Println(q)

	// Flow Controls are written in another package
	FlowControls()

	// More Types
	MoreTypes()

}
