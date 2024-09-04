package internal

import (
	"strings"
)

func reverse(list []string) []string {
	last := len(list) - 1
	for i := 0; i < len(list)/2; i++ {
		list[i], list[last-i] = list[last-i], list[i]
	}
	return list
}

func Tail(n int, input string) []string {
	lines := strings.Split(input, "\n")
	lines = reverse(lines)
	if n > len(lines) {
		return lines
	}
	return lines[:n]
}
