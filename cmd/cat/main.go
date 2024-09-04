package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	cmd "github.com/dohaelsawy/coreutils/internal"
	"log"
	"os"
	"strings"
)

var (
	ErrTxtNotFound = errors.New("the file you desire is not found :{")
	n              bool
	Scanner        *bufio.Scanner
)

func main() {
	flag.BoolVar(&n, "n", false, "numbered the lines of mentioned file")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 && args[0] != "-" {
		filePath := args[len(args)-1]
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(ErrTxtNotFound)
		}
		lines := cmd.Cat(n, string(data))
		for _, line := range lines {
			fmt.Println(line)
		}
		os.Exit(0)
	}
	args = os.Args[2:]
	lines := cmd.Cat(false, strings.Join(args, " "))
	fmt.Printf("%v\n", lines)
}
