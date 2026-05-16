package shell

import (
	"errors"
	"os/exec"
)

var errNotBuiltin = errors.New("not a builtin command")

func (s *Shell) execute(args []string) error {
	if len(args) == 0 {
		return nil
	}

	// Check for built-in commands.
	err := s.executeBuiltin(args)

	if err == nil {
		return nil
	}

	if !errors.Is(err, errNotBuiltin) {
		return err
	}

	return s.executeExternal(args)
}

func (s *Shell) executeExternal(args []string) error {
	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stdin = s.in
	cmd.Stdout = s.out
	cmd.Stderr = s.errOut

	// Execute the command and return the error
	return cmd.Run()
}
