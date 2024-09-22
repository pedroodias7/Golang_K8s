package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./hello-world <argument>\n")
		os.Exit()
	}


	fmt.Printf("Hello World\nos.Args: %v\n 1º Arguments: %v", args, args[1])

}