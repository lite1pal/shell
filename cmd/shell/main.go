package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Builtin func([]string) error

var builtins = map[string]Builtin{
	"cd":   cd,
	"exit": exitShell,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Add the indication to type
		fmt.Print("> ")

		// Read the keyboard input.
		input, err := readInput(reader)
		if errors.Is(err, io.EOF) {
			fmt.Println()
			return
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		args := parseInput(input)

		// Handle the execution of the input
		if err = execute(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func readInput(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\n')
}

func parseInput(input string) []string {
	return strings.Fields(input)
}

func execute(args []string) error {
	if len(args) == 0 {
		return nil
	}

	// Check for built-in commands.
	if err := executeBuiltin(args); err != nil {
		return err
	} else if !errors.Is(err, errNotBuiltin) {
		return err
	}

	return executeExternal(args)
}

var errNotBuiltin = errors.New("not a builtin command")

func executeBuiltin(args []string) error {
	fn, ok := builtins[args[0]]
	if !ok {
		return errNotBuiltin
	}

	return fn(args)
}

func executeExternal(args []string) error {
	// Prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command and return the error
	return cmd.Run()
}

func cd(args []string) error {
	if len(args) < 2 {
		return errors.New("cd: path required")
	}

	return os.Chdir(args[1])
}

func exitShell(args []string) error {
	os.Exit(0)
	return nil
}
