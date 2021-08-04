package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("{ Name = %v; Age = %v }", p.Name, p.Age)
}

func main() {
	p := Person{"Aarya", 21}

	fmt.Println(p)
}
