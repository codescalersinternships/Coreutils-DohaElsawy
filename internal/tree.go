package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound = errors.New("the file you desire is not found :{")
)

func TreeRecursive(a bool, l int, path, spaces string) error {
	if l != -2 && l < 0 {
		return nil
	}
	
	dirs, err := os.ReadDir(path)
	
	if err != nil {
		return ErrFileNotFound
	}
	
	if len(dirs) == 0 {
		return nil
	}
	
	if l != -2 {
		l--
	}
	
	for _, dir := range dirs {
		if dir.Name()[0] == '.' && !a {
			continue
		}
		
		if !dir.IsDir() {
			fmt.Printf("%s%s\n", spaces, dir.Name())
		}else {
			fmt.Printf("%s%s\n", spaces, dir.Name())
			err := TreeRecursive(a, l, filepath.Join(path, dir.Name()), spaces+"    ")
		
			if err != nil {
				return err
			}
		}
	}
	return nil
}
