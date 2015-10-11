package main

import (
    "os"
    "fmt"
)

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("Please specify the Yaml file for me to use.")
        return
    } else if len(args) > 1 {
        fmt.Println("Too many arguments! I can only use one file at a time.")
        return
    }
    filename := args[0]
    if err := Parse(filename); err != nil {
        panic(fmt.Errorf("%s\n\n%v","Oh dear, something went horribly wrong." , err))
    }
}
