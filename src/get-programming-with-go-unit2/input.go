package main

import (
    "fmt"
)

func main() {
	input := "3"
	var inputBool bool
	switch input {
	case "true":
	case "yes":
	case "1":
		inputBool = true
	case "false":
	case "no":
	case "0":
		inputBool = false
	default:
		fmt.Printf("convert fail for given '%v'\n", input)
	}
	result := fmt.Sprintf("convert result: %v", inputBool)
	fmt.Println(result)
}