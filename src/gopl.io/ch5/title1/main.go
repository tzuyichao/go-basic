package main

import (
    "fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [url]\n", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Printf("Getting title from %s\n", url)
	err := title(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func title(url string) error {
	fmt.Println("Title: ttttttt")
	return nil
}