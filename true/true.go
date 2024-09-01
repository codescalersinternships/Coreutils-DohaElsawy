package main

import (
	"os"
)


func True () bool{
	return true
}


func main() {
	flag:= True()
	if flag {
		os.Exit(0)
	}
}