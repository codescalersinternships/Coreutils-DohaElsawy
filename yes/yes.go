package main

import (
	"fmt"
	"os"
	"strings"
)



func Yes(output string) {
	for {
		fmt.Println(output)
	}
}

func main() {
	if len(os.Args) < 2 {
		Yes("y") // based on real experiment :} 
	}
	args := os.Args[1:]
	Yes(strings.Join(args," "))	
}