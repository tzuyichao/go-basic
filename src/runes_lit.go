package main

import (
    "fmt"
)

func main() {
    r1 := "™"
    fmt.Println(r1)
    fmt.Printf("%c\n", r1)

	r2 := '\x61'
	fmt.Println(r2)
	fmt.Printf("%c\n", r2)

	r3 := '\u2122'
	fmt.Println(r3)
	fmt.Printf("%c\n", r3)

	r4 := '\U00002122'
	fmt.Println(r4)
	fmt.Printf("%c\n", r4)
}