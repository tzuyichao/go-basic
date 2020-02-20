package main

import (
    "fmt"
    "strings"
)

// slice
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	// slice 
	planets := []string {" Venus     ", "Earth     ", "Mars"}
	fmt.Println(strings.Join(planets, "|"))
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}