package main

import (
	"fmt"
)

func main() {
	var s string = "Hello 🌞"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s, len(s))
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}