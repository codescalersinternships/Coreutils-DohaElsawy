package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
)

var (
	n int
)

func main() {
	flag.IntVar(&n, "n", 10, "last n lines from mentioned file")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {

		filePath := args[0]
		data, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatal(err)
		}

		lines := cmd.Tail(n, string(data))

		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
