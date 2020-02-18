package main

import (
    "fmt"
)

func main() {
	//cipherText := "CSOITEUIWUIZNSROCNKFD"
	//keyword := "GOLANG"
	a := 'A'
	k := 'E'
	c := 'X'
	t := 'T'
	fmt.Printf("%c\n", ((c-k)%26) + a)
	fmt.Printf("%c\n", ((t+k)%26) + a)
}