package main

import (
	"fmt"
	"time"
)

const (
    width = 80
    height = 15
)

func clear() {
	fmt.Print("\033c")
}

type Universe [][]bool

func NewUniverse() Universe {
	universe := make([][]bool, height)
	for i := 0; i < height; i++ {
		universe[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			universe[i][j] = false
		}
	}
	return universe
}

func (u Universe) show() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if u[i][j] {
				fmt.Print("T")
			} else {
				fmt.Print("F")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	clear()
	fmt.Printf("Game world (%dx%d)\n", width, height)
	time.Sleep(1000)
	clear()
	universe := NewUniverse()
	universe.show()
}