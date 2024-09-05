package main

import (
	"errors"
	"fmt"
	cmd "github.com/dohaelsawy/codescalers/coreutils/internal"
	"log"
	"os"
)

var (
	ErrNoEnv = errors.New("you don't have any env variables :{")
)

func main() {
	if len(os.Environ()) < 1 {
		log.Fatal(ErrNoEnv)
	}
	envList := cmd.Env()
	fmt.Println(envList)
}
