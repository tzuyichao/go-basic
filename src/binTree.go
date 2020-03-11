package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i:=0; i<2*n; i++ {
		temp := rand.Intn(n*2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, temp int) *Tree {
	if t == nil {
		return &Tree {nil, temp, nil}
	}
	if temp == t.Value {
		return t
	}
	if temp < t.Value {
		t.Left = insert(t.Left, temp)
		return t
	}
	t.Right = insert(t.Right, temp)
	return t
}

func main() {
	tree := create(10)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -20)
	traverse(tree)
	fmt.Println()
	fmt.Printf("Root value of tree is %d\n", tree.Value)
}