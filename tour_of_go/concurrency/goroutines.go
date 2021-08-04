package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 1; i <= 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	go say("Hello")
	say("World")
}
