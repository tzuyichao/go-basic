package main

import (
    "fmt"
)

func main() {
	var nowhere *int
	
	fmt.Println(nowhere)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(*nowhere)

	age := 15
	nowhere = &age

	if nowhere != nil {
		fmt.Println(*nowhere)
	}
}