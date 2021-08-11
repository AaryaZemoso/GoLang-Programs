package main

import "fmt"

func main() {

	messages := make(chan string)

	for i := 1; i <= 5; i++ {
		go func(x int) {
			for j := 1; j <= 10; j++ {
				fmt.Println(x, "->", j)
			}
			messages <- fmt.Sprintln("done: ", x)
		}(i)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-messages)
	}
}
