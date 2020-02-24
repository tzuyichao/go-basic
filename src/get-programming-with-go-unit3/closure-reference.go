package main

import "fmt"

type kalvin float64

func main() {
	var k kalvin = 294.0
	sensor := func() kalvin {
		return k
	}

	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
}