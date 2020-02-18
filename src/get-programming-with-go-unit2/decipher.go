package main

import (
    "fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	a := 'A'
	// k := 'E'
	// c := 'X'
	// fmt.Printf("%c", ((c-k)%26) + a)
	key_length := len(keyword)
	for idx, c := range cipherText {
		k_idx := idx % key_length
		fmt.Printf("%[1]c/%[1]v(%[1]T) map to keyword %[2]c/%[2]v(%[2]T)\n", c, keyword[k_idx])
		fmt.Printf("%c", ((c-int32(keyword[k_idx]))%26) + a)
	}
}