package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Shell struct {
	reader *bufio.Reader
	in     io.Reader
	out    io.Writer
	errOut io.Writer
}

func New(in io.Reader, out io.Writer, errOut io.Writer) *Shell {
	return &Shell{
		reader: bufio.NewReader(in),
		in:     in,
		out:    out,
		errOut: errOut,
	}
}

func (s *Shell) Run() error {
	for {
		if err := s.showInputIndicator(); err != nil {
			fmt.Fprintln(s.errOut, err)
			continue
		}

		input, err := s.readInput()
		if errors.Is(err, io.EOF) {
			fmt.Fprintln(s.out)
			return nil
		}

		if err != nil {
			fmt.Fprintln(s.errOut, err)
			continue
		}

		args := parseInput(input)

		if err := s.execute(args); err != nil {
			fmt.Fprintln(s.errOut, err)
		}
	}
}

func (s *Shell) showInputIndicator() error {
	workDir, err := os.Getwd()

	if err != nil {
		return err
	}

	fmt.Fprintf(s.out, "%s > ", workDir)

	return nil
}

func (s *Shell) readInput() (string, error) {
	return s.reader.ReadString('\n')
}
