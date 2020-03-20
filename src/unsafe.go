package main

import (
    "fmt"
    "unsafe"
)

func main() {
	var value int64 = 5
	var p1 *int64 = &value
	var p2 *int32 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)

	*p1 = 5434123412312431212 
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2)

	*p1 = 54341234
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2)
}