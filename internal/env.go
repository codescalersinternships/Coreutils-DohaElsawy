package internal

import (
	"os"
)

func Env() []string {
	var envOutput []string
	
	for _, env := range os.Environ() {
		envOutput = append(envOutput, env+"\n")
	}
	
	return envOutput
}
