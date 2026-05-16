package main

import (
	"fmt"
	"os"
	"shell/internal/shell"
)

func main() {
	sh := shell.New(os.Stdin, os.Stdout, os.Stderr)

	if err := sh.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
