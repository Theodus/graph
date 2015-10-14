package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please specify the JSON file for me to use.")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments!")
		return
	}
	
	filename := args[0]
	if err := Parse(filename); err != nil {
		panic(err)
	}
}
