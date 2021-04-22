package main

import (
    "fmt"
)

func main() {
    dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}

func double(number float64) float64 {
	return number * 2
}