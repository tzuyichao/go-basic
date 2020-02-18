package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Printf("%d bytes\n", len(question))
	for idx, c := range question {
		fmt.Printf("%v %c\n", idx, c)
	}
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size) 

	q1 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("%d bytes\n", len(q1))
	q2 := "¿"
	fmt.Printf("%d bytes\n", len(q2))
}