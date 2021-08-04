package main

import (
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS
	// No need of break statements
	// Exits once it finds appropriate case to be true
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	}

	// Switch with no condition is taken as switch true
	// This is a clean way to write long if else chains
	marks := 99
	switch {
	case marks > 90:
		fmt.Println("A grade")
	case marks > 80:
		fmt.Println("B grade")
	case marks > 70:
		fmt.Println("C grade")
	case marks > 60:
		fmt.Println("D grade")
	default:
		fmt.Println("F grade")
	}

}
