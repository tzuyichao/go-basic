package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if runtime.GOOS == "windows" {
					fmt.Println(fullPath)
				} else {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
						return
					}
				}
			}
		}
	}
}