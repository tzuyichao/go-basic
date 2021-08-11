package main

import (
	"fmt"
)

func remove(data []string, target string) []string {
	var result []string
	
	for i:=0; i<len(data); i++ {
		if data[i] != target {
			result = append(result, data[i])
		}
	}

	return result
}

func main() {
	var data = []string{
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
	}

	data = remove(data, "Bad")
	fmt.Println(data)
}