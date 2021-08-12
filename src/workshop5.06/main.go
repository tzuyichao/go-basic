package main

import (
	"fmt"
)

func main() {
	j := 9
	a := 1
	x := func(i int) int {
		return i * i
	}
	y := func(i int) int {
		return i * j
	}

	fmt.Printf("The square of %d is %d\n", j, x(j))
	fmt.Printf("The y(%d) is %d\n", a, y(a))
}