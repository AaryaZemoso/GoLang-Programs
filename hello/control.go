package main

import "fmt"

func FibonacciNthNumber(n int) (t int) {
	var f int = 0
	var s int = 1

	// For initialization; condition; updation {}
	for i := 2; i <= n; i++ {
		t = f + s
		f, s = s, t
	}

	return
}

func ReverseNumber(n int) int {

	temp := n
	reverse := 0

	for temp != 0 {
		reverse = reverse*10 + (temp % 10)
		temp /= 10
	}

	return reverse
}

func Sqrt(x float64) float64 {

	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func GradeMe(marks int) string {
	switch {
	case marks > 90:
		return "A"
	case marks > 80:
		return "B"
	case marks > 70:
		return "C"
	case marks > 60:
		return "D"
	default:
		return "F"
	}
}

func DeferTest() {
	defer fmt.Println("Defer Statement")

	fmt.Println("Statement 1")
	fmt.Println("Statement 2")
}

func FlowControls() {

	// Loops

	// For Loop
	fmt.Println("Fibonacci Number")
	fmt.Println(FibonacciNthNumber(10))

	// While Loop
	fmt.Println("Reversing a number using while loop")
	fmt.Println(ReverseNumber(13204))

	// If else statements
	// NOTE : writing else is new line is causing error?????
	fmt.Println("If else statemetns")
	r := 10
	if r < 20 {
		fmt.Println("Less than 20")
	} else {
		fmt.Println("Greater than 20")
	}

	// Excercise : Loops and functions
	fmt.Println("Excercise : Print sqrt of a given number")
	fmt.Println(Sqrt(99))

	// Switch as a long if-else statements
	fmt.Println("Switch Statement")
	fmt.Println(GradeMe(75))

	// Defer
	fmt.Println("Defer Test")
	DeferTest()
}
