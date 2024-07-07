package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}
	file := os.Args[1]
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	permissions := fileInfo.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permissions)
	fmt.Fprintf(os.Stdout, "Permissions: %s\n", permissionString)
}
