package main

import (
	"fmt"
	"os"
	"jogurto/commands"
)

func main() {
	var options string
	var packages []string
	if len(os.Args) < 2 {
		fmt.Printf("%s: A yaourt-like interface and toolkit for Debian.", os.Args[0])
	} else {
		options = os.Args[1]
		packages = os.Args[2:]
	}
	run := commands.Map[options]
	if run != nil {
		err := run(packages)
		if err != nil {
			fmt.Printf("An error occurred: %s", err.Error())
		}
	} else {
		fmt.Println("Command not defined")
	}
}