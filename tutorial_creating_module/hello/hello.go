package main

import (
	"fmt"
	"gettingstarted/greetings"
	// "rsc.io/quote"
)

func main() {
	message, ok := greetings.Hello("Aarya")

	if ok == nil {
		fmt.Println(message)
	} else {
		fmt.Println(ok.Error())
	}
	// fmt.Println(quote.Go())
}
