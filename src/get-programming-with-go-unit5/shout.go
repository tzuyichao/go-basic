package main

import (
	"fmt"
	"strings"
)

type talk interface {
	talk() string
}

func shout(t talk) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct {}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type starship struct {
	laser
}

func main() {
	m := martian{}
	shout(m)

	l := laser(4)
	shout(l)

	s := starship{
		laser(3),
	}
	shout(s)
}