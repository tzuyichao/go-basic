package main

import (
	"fmt"
	"math"
)

// from Go Programming by Example
func main() {
	fmt.Println("Circle Area caculation")
	fmt.Print("Enter a radius value: ")
	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Circle area with radius %.2f = %.2f\n", radius, area)
}