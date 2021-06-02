package main

import (
	"fmt"
)

func powerSeries(a int) (int, int) {
	s := a*a
	return s, s*a
}

func main() {
	var square int
	var cube int

	square, cube = powerSeries(3)
	
	fmt.Println("Square: ", square, ", Cube: ", cube)
}