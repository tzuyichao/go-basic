package main

import (
    "fmt"
)

func main() {
    text := "WEDIGYOULUVTHEGOPHERS"
	keyword := "GOLANG"
	a := 'A'
	key_length := len(keyword)
	for idx, c := range text {
		ch := c
		k_idx := idx % key_length
		fmt.Printf("%c", ((ch+int32(keyword[k_idx]))%26) + a)
	}
}