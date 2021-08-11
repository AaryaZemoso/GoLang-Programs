package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second)
		c1 <- "three"
	}()
	go func() {
		time.Sleep(time.Second)
		c1 <- "four"
	}()
	go func() {
		time.Sleep(time.Second)
		c1 <- "five"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- "six"
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- "seven"
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- "eight"
	}()

	for i := 0; i < 8; i++ {
		select {
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg1 := <-c1:
			fmt.Println(msg1)

		}
	}
}
