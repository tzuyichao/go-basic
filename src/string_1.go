package main

import (
	"fmt"
)

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func main() {
	var a ArrayValue
	a = []string{"abc", "def"}
	fmt.Printf(a.String())
}