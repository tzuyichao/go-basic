package main

import (
    "fmt"
)

type turtle struct {
	x, y int
}

func up(t *turtle) {
	t.y++
}

func down(t *turtle) {
	t.y--
}

func right(t *turtle) {
	t.x++
}

func left(t *turtle) {
	t.x--
}

func (t turtle) inspect() {
	fmt.Printf("(x: %d, y: %d)", t.x, t.y)
}

func main() {
	t := &turtle {10, 10}
	up(t)
	up(t)
	left(t)
	down(t)
	right(t)
	t.inspect()
}