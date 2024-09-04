package internal

import (
	"strings"
)

func Head(n int, input string) []string {
	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	if n > len(lines) {
		return lines
	}
	return lines[:n]
}
