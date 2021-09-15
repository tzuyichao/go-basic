package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a file")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	var (
		b = make([]byte, 16)
	)
	for n:=0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Print(string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError:", err)
	}
}