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

	insight := location{4.5, 135.9}
	fmt.Println(insight)
}