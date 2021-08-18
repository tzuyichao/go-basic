package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File Content:")
	fmt.Println(string(content))
}