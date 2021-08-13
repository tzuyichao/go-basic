package main

import (
	"fmt"
)

func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is string\n", v)
		case bool:
			fmt.Printf("%v is bool\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
}

type cat struct {
	name string
}

func main() {
	c := cat{name: "Bob"}
	i := []interface{}{43, "The book club", true, c}
	typeExample(i)
}