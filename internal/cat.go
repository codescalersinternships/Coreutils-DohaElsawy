package internal

import (
	"fmt"
	"strings"
)

func Cat(n bool, input string) []string {
	lines := strings.Split(input, "\n")

	if n {
		for i := range lines {
			lines[i] = fmt.Sprintf("%d    %s", i+1, lines[i])
		}
		return lines
	}

	return lines
}
