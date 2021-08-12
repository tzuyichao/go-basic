package main

import (
	"fmt"
)

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))

	// anonymous function
	func() {
		fmt.Println("Greeting...")
	}()

	f := func() {
		fmt.Println("Executing an anonymous function using a variable")
	}
	fmt.Println("Line after anonymous function")
	f()
}

func sum(i ...int) int {
	total := 0
	for _, num := range i {
		total += num
	}
	return total
}