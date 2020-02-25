package main

import (
    "fmt"
)

func main() {
    answer := 42
	fmt.Println(&answer)
	fmt.Println(*&answer)
	
	address := &answer
	fmt.Println(*address)
	fmt.Printf("address is a %T\n", address)

	canada := "canada"
	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Printf("%v\t%v", home, *home)
}