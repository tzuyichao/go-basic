package main

import (
	"fmt"
)

func main() {
    temperatures := []float64 {
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	fmt.Println(temperatures)

	fmt.Println("Frequency")
	for t, count := range frequency {
		fmt.Printf("%5.1f: %02d\n", t, count)
	}
}