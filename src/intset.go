package main

import (
    "fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word] & (1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1<<bit
}

func main() {
	var a IntSet
	a.Add(1)
	a.Add(9)
	fmt.Printf("%v\n", a)
	fmt.Printf("has 1? %v\n", a.Has(1))
	fmt.Printf("has 9? %v\n", a.Has(9))
	fmt.Printf("has 4? %v\n", a.Has(4))
}