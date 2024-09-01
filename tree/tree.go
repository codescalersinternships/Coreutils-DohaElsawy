package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)


type ErrTypes string

const (
	ErrFileNotFound = ErrTypes("the file you desire is not found :{")
)


func (e ErrTypes) Error() string {
	return string(e)
}

func TreeRecursive(a bool, l int, path,spaces string){
	if l != -2  && l < 0{
		return
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(dirs ,ErrFileNotFound)
		return
	}
	if len(dirs) == 0 {
		return 
	}
	if l != -2{
		l--
	}
	for _, dir := range dirs {
		if dir.IsDir(){
			if dir.Name()[0] == '.' && !a {
				continue
			} 
			fmt.Printf("%s%s\n", spaces ,dir.Name())
			TreeRecursive(a,l,filepath.Join(path, dir.Name()), spaces+"    ")
		}else {
			fmt.Printf("%s%s\n", spaces, dir.Name())
		}
	}
}

func main() {
	var l int
	var a bool
	flag.IntVar(&l,"L", -2, "last n lines from mentioned file")
	flag.BoolVar(&a,"a",false, "display hidden files")
	flag.Parse()
	// fmt.Println(a)
	args := flag.Args()
	if len(args) > 0 {
		fileName := args[len(args)-1]
		TreeRecursive(a,l,fileName,"")
	}
}
