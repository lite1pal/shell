package shell

import "strings"

func parseInput(input string) []string {
	return strings.Fields(input)
}
