package main

import (
	"fmt"
)

func test() (*int) {
	x := 10
	return &x
}

func main() {
	res := test()
	fmt.Println(*res)
}