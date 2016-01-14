package main

import (
	"fmt"
	"importingFns/arithmetic"
)

func main() {
	a := 5
	b := 7
	sum := arithmetic.AddInts(a, b)
	fmt.Println(a, "+", b, "=", sum)
}
