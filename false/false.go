package main

import (
	"os"
)

func False () (bool){
	return true
}

func main() {
	flag := False()
	if flag {
		os.Exit(1)
	}
}