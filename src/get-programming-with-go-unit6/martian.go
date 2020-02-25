package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct {}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	m := martian{}
	
	shout(m)
	shout(&m)

	l := laser(3)
	shout(&l)
	// laser does not implement talker (talk method has pointer receiver)
	// shout(l)
}