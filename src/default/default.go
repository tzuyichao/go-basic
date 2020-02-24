package main

import "fmt"

func main() {
	var price float64
	fmt.Println(price)
	fmt.Println("price: ", price)
	
	var third = 0.33
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%05.2f\n", third)

	third = 1.0/3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
}