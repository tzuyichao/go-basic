package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"path/filepath"
)

func ScanDirectory(path string) error {
	fmt.Println("Directory:", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := ScanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("File:", filePath)
		}
	}
	return nil
}

func main() {
    err := ScanDirectory("my_directory")
    if err != nil {
		log.Fatal(err)
	}
}