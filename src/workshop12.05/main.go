package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFile(filename string) {
	fileinfo, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%v: file not exist!\n\n", filename)
			return
		}
	}
	fmt.Printf("Filename: %s\n是否目錄: %t\n修改時間: %v\n權限: %v\n大小: %d\n\n", fileinfo.Name(), fileinfo.IsDir(), fileinfo.ModTime(), fileinfo.Mode(), fileinfo.Size())
}

func main() {
	checkFile("junk.txt")
	checkFile("go.mod")
	checkFile(".")
}