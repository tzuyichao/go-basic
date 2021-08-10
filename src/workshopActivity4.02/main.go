package main

import (
	"fmt"
	"os"
)

func main() {
	var data map[string]string
	data = make(map[string]string)

	data["305"] = "Sue"
	data["204"] = "Bob"
	data["631"] = "Jake"
	data["073"] = "Tracy"

	if len(os.Args) != 2 {
		fmt.Println("參數不足")
		os.Exit(1)
	}
	key := os.Args[1]
	name, ok := data[key]
	if ok {
		fmt.Println("Hi, ", name)
	} else {
		fmt.Println("Not found: ", key)
	}
}