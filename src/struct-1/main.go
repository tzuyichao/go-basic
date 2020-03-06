package main

import (
    "fmt"
)

func main() {
	// anonymous struct and struct literal
    data := struct {
		Name string
	} {
		Name: "text1",
	}

	fmt.Println(data)
}