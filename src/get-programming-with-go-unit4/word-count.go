package main

import (
    "fmt"
    "strings"
)

func main() {
	str := "Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode. IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space."
	str = strings.ToLower(str)
	terms := strings.Fields(str)
	wc := make(map[string]int32)
	for _, term := range terms {
		term = strings.Trim(term, " .,!?")
		wc[term]++
	}
	fmt.Println(str)
	for term, num := range wc {
		fmt.Printf("%s: %d\n", term, num)
	}
}