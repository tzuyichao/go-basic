package main

import (
    "fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	// cipherText := "LXFOPVEFRNHR"
	// keyword := "LEMON"
	a := int32('A')
	// k := 'E'
	// c := 'X'
	// fmt.Printf("%c", ((c-k)%26) + a)
	key_length := len(keyword)
	for idx, c := range cipherText {
		k_idx := idx % key_length
		kc := int32(keyword[k_idx])
		var e int32
		if c-kc > 0 {
			e = (c-kc)%26+a
		} else {
			e = (c-kc+26)%26+a
		}
		fmt.Printf("%d %v %c %v %c %v %c\n", idx, c, c, kc, kc, e, e)
	}
}