package main

import (
	"fmt"
)

type item struct {
	name string
	val int
}

func main() {
	evenVals := []item{
		item{"2", 2},
		item{"4", 4},
		item{"6", 6},
		item{"8", 8},
		item{"10", 10},
		item{"12", 12},
	}
	for _, v := range evenVals {
		v.val *= 2
		fmt.Println(v)
	}
	fmt.Println(evenVals)
}