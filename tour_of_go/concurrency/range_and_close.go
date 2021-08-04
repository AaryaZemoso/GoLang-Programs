package main

import "fmt"

func fib(n int, c chan int) {
	f, s := 0, 1
	for i := 0; i < n; i++ {
		c <- f
		f, s = s, f+s
	}
	close(c)
}

func main() {

	c := make(chan int)
	go fib(10, c)

	for i := range c {
		fmt.Println(i)
	}

}
