package internal

import (
	"fmt"
	"strings"
)

func Cat(n bool, input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if n {
		for i := range lines {
			lines[i] = fmt.Sprintf("%d %s", i+1, strings.TrimSpace(lines[i]))
		}
	} else {
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}
	}
	return lines
}
