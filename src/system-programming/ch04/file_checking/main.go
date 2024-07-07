package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], " <file>")
		os.Exit(1)
	}
	file := os.Args[1]
	info, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "File name: %s\n", info.Name())
	fmt.Fprintf(os.Stdout, "File size: %d\n", info.Size())
	fmt.Fprintf(os.Stdout, "File permission: %s\n", info.Mode())
	fmt.Fprintf(os.Stdout, "Last modified: %s\n", info.ModTime())
}
