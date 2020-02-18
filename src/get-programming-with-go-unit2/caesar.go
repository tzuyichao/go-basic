package main

import "fmt"

func main() {
    c := 'a'
    c = c + 3
	fmt.Printf("%c %[1]v\n", c)
	
	c = 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c\n", c)

	c = 'A'
	fmt.Printf("A: %v\n", c)

	for i := 'A'; i <= 'z'; i++ {
		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			fmt.Printf("%c ", i)
		}
	}
}