package main

import (
	"errors"
	"flag"
	"fmt"
	cmd "github.com/dohaelsawy/coreutils/internal"
	"log"
	"os"
)

var (
	ErrTxtNotFound = errors.New("the file you desire is not found :{")
	n              int
)

func main() {
	flag.IntVar(&n, "n", 10, "last n lines from mentioned file")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		filePath := args[0]
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(ErrTxtNotFound)
		}
		lines := cmd.Tail(n, string(data))
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
