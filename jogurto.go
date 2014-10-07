package main

import (
	"fmt"
	"github.com/lapingvino/jogurto/commands"
	"os"
)

func main() {
	var options string
	var packages []string
	if len(os.Args) < 2 {
		fmt.Printf("%s: A yaourt-like interface and toolkit for Debian.\n\n", os.Args[0])
	} else {
		options = os.Args[1]
		packages = os.Args[2:]
	}
	run := commands.Map[options]
	if run != nil {
		err := run(packages)
		if err != nil {
			fmt.Println("An error occurred:", err.Error())
			os.Exit(1)
		}
	} else {
		if options != "" {
			fmt.Println("Option", options, "not implemented\n")
			os.Exit(2)
		}
		commands.Help(nil)
	}
}
