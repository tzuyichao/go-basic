package main

import "fmt"

var f = func() {
    fmt.Println("Dress up for the masquerade.")
}

func main() {
	f()
	g := func(message string) {
		fmt.Printf("Message: %v\n", message)
	}
	g("Go to the party.")
	func(message string) {
		fmt.Printf("Message: %v\n", message)
	}("Let's shake our body.")
}