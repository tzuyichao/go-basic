package main

import (
    "fmt"
    "math/rand"
)

func main() {
	// print header
	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")
	// print ticket
    for i := 0; i < 10; i++ {
		price := rand.Intn(70) + 30
		days := rand.Intn(30) + 20
		triptype := "One-way"
		if price > 45 {
			triptype = "Round-trip"
		}
		var spaceline = "Virgin Galactic"
		switch num := rand.Intn(3); num {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Space Adventures"
		}
		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, days, triptype, price)
	}
}