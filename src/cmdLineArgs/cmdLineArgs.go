package main

import (
	"fmt"
	"os"
)

func main() {
	cmdLineArgs := os.Args[1:] // skip program name
	if len(cmdLineArgs) != 2 {
		fmt.Println("Usage: cmdLineArgs <arg1> <arg2>")
		os.Exit(-1)
	}
	arg1 := cmdLineArgs[0]
	arg2 := cmdLineArgs[1]
	fmt.Println("arg1:", arg1, "arg2:", arg2)
}
