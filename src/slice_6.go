package main

import (
	"fmt"
)

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	for i := 1; i<6; i++ {
		x = append(x, i * 10)
		fmt.Println(x, len(x), cap(x))
	}
}