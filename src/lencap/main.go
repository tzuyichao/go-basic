package main

import (
	"fmt"
)

func inspect(x []int, s int) []int {
	res := x
	res = append(res, s)
	fmt.Println(res, len(res), cap(res))
	return res
}

func main() {
	var x []int

	for _, elem := range [5]int{10, 20, 30, 40, 50} {
		x = inspect(x, elem)
	}
}