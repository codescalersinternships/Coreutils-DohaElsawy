package main

import (
	"fmt"
	"os"
)

type ErrTypes string

const (
	ErrNoEnv = ErrTypes("you don't have any env varialbles :{")
)

func (e ErrTypes) Error() string {
	return string(e)
}

func Env() error{
	if len(os.Environ()) < 1 {
		return ErrNoEnv
	}
	for _,env := range os.Environ() {
		fmt.Println(env)
	} 
	return nil
}

func main() {
	err := Env()
	if err == ErrNoEnv {
		fmt.Println(err)
	}
}