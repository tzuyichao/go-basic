package main

import (
    "fmt"
)

type location struct {
	lat, long float64
}

func (l location) String() string {
	fmt.Println("String() called")
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct {
	d, m, s float64
	h rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%d°%d'%.1f\" %c", int(c.d), int(c.m), c.s, c.h)
}

func main() {
	curiosity := location {-4.5895, 137.4417}
	fmt.Println(curiosity)

	ep1 := coordinate {4.0, 30.0, 0.0, 'N'}
	ep2 := coordinate {135.0, 54.0, 0.0, 'E'}
    fmt.Printf("Elysium Planitia is at %v, %v", ep1, ep2)
}