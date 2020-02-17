package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang rune = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	smile := '😃'
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)
}