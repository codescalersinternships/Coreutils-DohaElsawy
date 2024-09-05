package main

import (
	"flag"
	"fmt"
	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
	"strings"
)

var (
	n bool
)

func main() {

	flag.BoolVar(&n, "n", false, "omit the trailing newline")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		output := cmd.Echo(n, strings.Join(args, " "))
		fmt.Println(output)
		return
	} 
	
	fmt.Println(" ")
}
