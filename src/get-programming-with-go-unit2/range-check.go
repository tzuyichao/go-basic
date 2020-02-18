package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32768
	var h int16 = int16(0)
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("exceed")
	} else {
		h = int16(bh)
		fmt.Println(h)
	}
}