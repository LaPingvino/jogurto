package commands

import "fmt"
import "os"

func Help(command []string) error {
	_, err := fmt.Println("Usage:", os.Args[0], `option [packages]

Where options is one of:
    --help, -h: show this help
    -S: install packages
    -Sy: update package database [and install]
    -Syu: update && upgrade && install
    -Syyu: update && dist-upgrade && install
    -Ss: search in installable packages
    -R: remove packages
    -Rn: purge packages
    -Rs: remove packages, then autoremove dependencies
    -Q: show all installed packages
`)

	return err
}
