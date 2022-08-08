package main

import (
	"fmt"
)

type Flag int

func main() {
	var i interface{} = 3
	f, ok := i.(Flag)
	if !ok {
		fmt.Println("not a flag")
		return
	}
	fmt.Println(f)
}