package main

import "fmt"

func main() {
	age := 44
	marsDays := 687
	earthDays := 365.2425
	// invalid operation: age * earthDays (mismatched types int and float64)
	fmt.Println("I am", age*earthDays/marsDays + "years old on Mars.")
}