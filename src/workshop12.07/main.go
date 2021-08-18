package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File Content:")
	fmt.Println(string(content))
}