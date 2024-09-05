package main

import (
	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		cmd.Yes("y") // based on real experiment :}
	}
	args := os.Args[1:]
	cmd.Yes(strings.Join(args, " "))
}
