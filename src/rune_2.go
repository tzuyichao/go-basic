package main

import (
	"fmt"
)

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Println()

	for i := 0; i < len(username); i++ {
		fmt.Printf("%x ", username[i])
	}
	fmt.Println()

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}
}