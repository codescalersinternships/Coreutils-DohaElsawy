package internal

import (
	"strings"
)

func WordCount(input string) int {
	return len(strings.Fields(input))
}
func CharCount(input string) int {
	return len(input)
}
func LineCount(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return len(lines)
}

func WC(l, w, c bool, input string) (int, int, int) {
	return LineCount(input), WordCount(input), CharCount(input)
}
