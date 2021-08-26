package main

import (
	"fmt"
)

const maxUint = ^uint(0)

func IncUint(counter uint) uint {
	if counter == maxUint {
		panic("uint overflow")
	}
	return counter + 1
}

func main() {
	fmt.Println("maxUint:", maxUint)
	var x uint = maxUint - 1
	x = IncUint(x)
	fmt.Println("x = x + 1: ", x)
	x = IncUint(x)
	fmt.Println("x = x + 1 again done. ", x)
}