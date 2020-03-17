package main

import (
	"fmt"
)

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}

	return &myStructure{n, s, h}
}

func retStruct(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	m1 := createStruct("M1", "M1s", 310)
	fmt.Println(m1)
	fmt.Println(*m1)
	fmt.Println(m1.Name)
	fmt.Println((*m1).Name)
	m2 := retStruct("M2", "M2s", 399)
	fmt.Println(m2)

	p1 := myStructure{}
	pnew := new(myStructure)
	fmt.Println(p1)
	fmt.Println(pnew)
	fmt.Println("slice -----")
	sp := []myStructure{}
	spnew := new([]myStructure)
	fmt.Println(sp)
	fmt.Println(spnew)
}
