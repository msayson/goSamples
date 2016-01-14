package main

import (
	"fmt"
)

func main() {
	//Explicit type declaration
	var name string
	name = "Fred"

	//Implicit type inference
	greeting := "Hello"
	fmt.Println(greeting, name)
}
