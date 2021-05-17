package main

import (
	"fmt"
)

func main() {
	d := [4]int{5, 6, 7, 8}
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
	num = copy(d[:], x)
	fmt.Println(d, num)
}