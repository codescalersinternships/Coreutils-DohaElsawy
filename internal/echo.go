package internal

import (
	"strings"
)

func Echo(n bool, input string) string {
	if n {
		return strings.Replace(input, "\r\n", "", -1)
	}
	return input
}
