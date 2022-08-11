package main

import (
	"fmt"
)

func main() {
	msg := "π = 3.14159265358..."
	fmt.Printf("%T\n", msg[0])
	for _, c := range msg {
		fmt.Printf("%T\n", c)
		break
	}
}