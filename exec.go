package mcore

import (
	"os/exec"
	"path/filepath"
)

// NewCommand return new exec cmd
// args[0] cmd name
// args[1:] cmd args
func NewCommand(args []string) *exec.Cmd {
	if len(args) == 0 {
		return nil
	}
	name := args[0]
	cmd := &exec.Cmd{
		Path: name,
		Args: args,
	}

	if filepath.Base(name) == name {
		if lp, err := exec.LookPath(name); err != nil {
			//cmd.lookPathErr = err
		} else {
			cmd.Path = lp
		}
	}
	return cmd
}

// NewCommandFromString
func NewCommandFromString(cmd string) *exec.Cmd {
	return NewCommand(ParseStringToArgs(cmd))
}
