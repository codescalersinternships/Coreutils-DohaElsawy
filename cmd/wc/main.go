package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
)

var (
	w, l, c bool
)

func main() {

	flag.BoolVar(&w, "w", false, "last n words from mentioned file")
	flag.BoolVar(&l, "l", false, "last n lines from mentioned file")
	flag.BoolVar(&c, "c", false, "last n characters from mentioned file")

	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		filePath := args[0]
		data, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatal(err)
		}

		lines, words, chars := cmd.WC(l, w, c, string(data))

		if !l && !w && !c {
			fmt.Println(lines, words, chars)
			return
		}
		if l {
			fmt.Println(lines)
		}
		if w {
			fmt.Println(words)
		}
		if c {
			fmt.Println(chars)
		}

	}
}
