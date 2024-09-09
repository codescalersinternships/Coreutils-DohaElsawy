package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
)

var (
	n       bool
	Scanner *bufio.Scanner
)

func main() {

	flag.BoolVar(&n, "n", false, "numbered the lines of mentioned file")
	flag.Parse()

	args := flag.Args()

	scanner := bufio.NewScanner(os.Stdin)

	if len(args) > 0 && args[0] != "-" {

		filePath := args[len(args)-1]
		data, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatal(err)
		}

		lines := cmd.Cat(n, string(data))

		for _, line := range lines {
			fmt.Println(line)
		}

		return
	}

	for scanner.Scan() {
		lines := cmd.Cat(n, scanner.Text())
		fmt.Println(lines)
	}
}
