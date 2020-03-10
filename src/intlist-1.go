package main

import (
    "fmt"
)

type IntList struct {
	Value int
	Tail *IntList
}

func (intList *IntList) Sum() int {
	if intList == nil {
		return 0
	}
	return intList.Value + intList.Tail.Sum()
}

func (intList *IntList) Insert(value int) *IntList {
	if intList == nil {
		return &IntList {value, nil}
	}
	if intList.Tail == nil {
		intList.Tail = &IntList {value, nil}
	} else {
		intList.Tail.Insert(value)
	}
	return intList
}

func main() {
	test1 := &IntList {10, nil}
	var test2 IntList
	fmt.Printf("test2 sum=%d\n", test2.Sum())
	fmt.Printf("test1 sum=%d\n", test1.Sum())
	test1.Insert(20)
	fmt.Printf("test1 sum=%d\n", test1.Sum())
}