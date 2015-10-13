package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
  cmd("go", []string{"install", "graph"})
  fmt.Println("binary palaced in bin/")
}

func cmd(cmdName string, cmdArgs []string) {
  var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "error running command: ", err)
		os.Exit(1)
	}
  fmt.Println(string(cmdOut))
}
