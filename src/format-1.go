package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}
