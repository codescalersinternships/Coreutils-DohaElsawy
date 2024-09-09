package internal

import (
	"strings"
)

func Tail(n int, input string) []string {
	lines := strings.Split(input, "\n")

	if n > len(lines) {
		return lines
	}

	return lines[len(lines)-n:]
}
