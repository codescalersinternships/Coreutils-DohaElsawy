package main

import (
	"errors"
	"flag"
	cmd "github.com/dohaelsawy/coreutils/internal"
	"log"
)

var (
	l               int
	a               bool
	ErrFileNotFound = errors.New("the file you desire is not found :{")
)

func main() {

	flag.IntVar(&l, "L", -2, "last n lines from mentioned file")
	flag.BoolVar(&a, "a", false, "display hidden files")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		fileName := args[len(args)-1]
		err := cmd.TreeRecursive(a, l, fileName, "")
		if err != nil {
			log.Fatal(err)
		}
	}
}
