package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Printf("%d bytes\n", len(question))
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size) 
}