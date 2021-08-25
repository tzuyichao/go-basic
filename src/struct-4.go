package main

import (
	"fmt"
)

func main() {
	var person struct {
		name string
		age int
		pet string
	}

	person.name = "Bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "Dog",
	}

	fmt.Println(person)
	fmt.Println(pet)
}