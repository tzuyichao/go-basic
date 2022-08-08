package main

import (
	"fmt"
)

const (
	Read = 1 << iota
	Write
	Execute
)

func main() {
	mask := Read | Execute
	
	if mask & Execute == 0 {
		fmt.Println("Can't Execute")
	} else {
		fmt.Println("Can Execute")
	}
	
	if mask & Write == 0 {
		fmt.Println("Can't Write")
	} else {
		fmt.Println("Can Write")
	}
}