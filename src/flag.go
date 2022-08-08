package main

import (
	"fmt"
)

type Flag int

func main() {
	var i interface{} = 3
	f := i.(Flag)
	fmt.Println(f)
}