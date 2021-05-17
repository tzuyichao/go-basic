package main

import (
    "fmt"
)

func main() {
    type person struct {
		name string
		age int
		pet string
	}

	var fred person
	bob := person{}
	julia := person{
		"julia",
		40,
		"cat",
	}
	beth := person{
		age: 30,
		name: "Beth",
	}

	fmt.Println("Fred:", fred)
	fmt.Println("Bob:", bob)
	fmt.Println("Julia:", julia)
	fmt.Println("Beth:", beth)
}