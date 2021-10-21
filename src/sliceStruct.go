package main

import (
	"fmt"
	"strconv"
)

type record struct {
	Field1 int 
	Field2 string
}

func main() {
	S := []record{}
	fmt.Println(S)
	fmt.Println(len(S))
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		S = append(S, temp)
	}
	fmt.Println(S)
	fmt.Println(len(S))
}