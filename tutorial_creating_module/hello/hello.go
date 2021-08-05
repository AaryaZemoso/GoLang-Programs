package main

import (
	"fmt"
	"gettingstarted/greetings"
	"os"
	// "rsc.io/quote"
)

func main() {

	args := os.Args
	if len(args) < 1 {
		// fmt.Println("Command line arguments : ", args[1:])
		fmt.Println("Provide arguments")
		return
	}

	names := args[1:]

	message, ok := greetings.Hello(names[0])

	if ok == nil {
		fmt.Println(message)
	} else {
		fmt.Println(ok.Error())
	}
	// fmt.Println(quote.Go())

	greetingsMessages, err := greetings.Hellos(names)
	if err == nil {
		for k, v := range greetingsMessages {
			fmt.Println(k, "->", v)
		}
	} else {
		fmt.Println(err)
	}

}
