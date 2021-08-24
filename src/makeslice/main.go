package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 10)
	fmt.Println(x, len(x), cap(x))
	var data []int
	fmt.Println("data == nil:", data == nil)
	var data2 = []int{}
	fmt.Println("data2 == nil:", data2 == nil)
	// append一定會增加slice的長度
	y := make([]int, 5, 10)
	y = append(y, 10)
	fmt.Println(y, len(y), cap(y))
}