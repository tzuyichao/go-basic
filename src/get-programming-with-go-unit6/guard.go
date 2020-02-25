package main

import (
    "fmt"
)

type person struct {
	age int
}

func (p *person) birthday_without_guard() {
	p.age++
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	var noperson *person
	noperson.birthday()
	fmt.Println(noperson)
	// panic: runtime error: invalid memory address or nil pointer dereference
	//noperson.birthday_without_guard()

	var fn func(a, b int) int
	fmt.Println(fn == nil)
}