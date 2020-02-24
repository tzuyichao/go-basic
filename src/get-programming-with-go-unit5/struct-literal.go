package main

import (
    "fmt"
)

func main() {
    type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)
}