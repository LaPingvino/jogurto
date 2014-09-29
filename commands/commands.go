package commands

import "os/exec"
import "os"

type Command func (packages []string) error

func doAll(fn ...Command) Command {
	return func (packages []string) error {
		var err error
		for i := range fn {
			err = fn[i](packages)
			if err != nil { return err }
		}
		return nil
	}
}

func run(command string, arg string, args []string) error {
	var cmdargs []string
	if arg == "" {
		cmdargs = args
	} else {
		cmdargs = append([]string{arg}, args...)
	}
	cmd := exec.Command(command, cmdargs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var Map = map[string]Command {
	"-S": Install,
	"-Sy": doAll(Update, Install),
	"-Syu": doAll(Update, Upgrade, Install),
	"-Syyu": doAll(Update, DistUpgrade, Install),
	"-Ss": Search,
} 