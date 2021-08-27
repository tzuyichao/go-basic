package main

import (
	"fmt"
)

func main() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, token := range sample {
			fmt.Println(i, string(token))
		}
		fmt.Println()
	}
}