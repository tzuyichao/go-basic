package main

import (
	"fmt"
)

func defineArray() [10]int {
	var res [10]int
	return res
}

func main() {
	fmt.Printf("%#v\n", defineArray())
}