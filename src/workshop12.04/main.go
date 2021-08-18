package main

import (
	"os"
)

func main() {
	message := "Hello Golang!"
	err := os.WriteFile("test.txt", []byte(message), 0644)
	if err != nil {
		panic(err)
	}
}