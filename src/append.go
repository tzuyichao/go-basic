package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3}
	fmt.Printf("a: %T\n", a)
    b := append(a[:1], 10)
	fmt.Printf("b: %T\n", b)
    fmt.Printf("a=%v, b=%v", a, b)
}