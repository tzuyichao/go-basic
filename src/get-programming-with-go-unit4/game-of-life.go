package main

import (
    "fmt"
)

const (
    width = 80
    height = 15
)

type Universe [][]bool

func main() {
	fmt.Printf("Game world (%dx%d)\n", width, height)
}