package shell

import (
	"errors"
	"os"
)

type Builtin func(*Shell, []string) error

var builtins = map[string]Builtin{
	"cd":   cd,
	"exit": exitShell,
}

func (s *Shell) executeBuiltin(args []string) error {
	fn, ok := builtins[args[0]]
	if !ok {
		return errNotBuiltin
	}

	return fn(s, args)
}

func cd(s *Shell, args []string) error {
	if len(args) < 2 {
		return errors.New("cd: path required")
	}

	return os.Chdir(args[1])
}

func exitShell(s *Shell, args []string) error {
	os.Exit(0)
	return nil
}
