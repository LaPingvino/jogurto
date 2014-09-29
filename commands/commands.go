package commands

import "os/exec"
import "os"

type Command func (packages []string) error

// doAll permits to perform commands one after another
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

// run does everything to run a command as if you typed it
// yourself. The arg part can be omitted giving it an empty
// string. It's there mostly to make the apt commands work
// as they have the actual command as the first arg.
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

// The options that you can use with jogurto
// You can put functions here, or string functions together
// with a higher order function like doAll.
var Map = map[string]Command {
	"-S": Install,
	"-Sy": doAll(Update, Install),
	"-Syu": doAll(Update, Upgrade, Install),
	"-Syyu": doAll(Update, DistUpgrade, Install),
	"-Ss": Search,
	"-R": Remove,
	"-Rn": Purge,
} 