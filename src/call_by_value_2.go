package main

import (
	"log"
	"fmt"
)

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for idx, val := range s {
		s[idx] = val * 2
	}
	s = append(s, 10)
	log.Println("in mod: ", s)
}

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)

	var s2 = make([]int, 0, 10)
	s2 = append(s2, 1, 2)
	modSlice(s2)
	fmt.Println(s2)
}