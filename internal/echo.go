package internal

import (
	"strings"
)

func Echo(n bool, input []string) string {
	ans := strings.Join(input, " ")

	if n {
		return ans
	}
	ans += "\n"
	return ans
}
