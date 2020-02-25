package main

import (
	"fmt"
	"math"
)

type rover struct {
	gps
}

type gps struct {
	current location
	destination location
	world
}

func (g gps) message() string {
	return fmt.Sprintf("distance to %s is %f km", g.destination.name, g.distance(g.current, g.destination))
}

type world struct {
	radius float64
}

func rad(deg float64) float64 {
    return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
    s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type location struct {
	name string
	lat float64
	long float64
}

func main() {
	fmt.Println("GPS Homework")
	rover := rover {
		gps {
			current: location {"spirit", -14.5684, 175.472636},
			destination: location {"opportunity", -1.9462, 354.4734},
			world: world {radius: 3389.5},
		},
	}
	fmt.Println(rover.distance(rover.current, rover.destination))
	fmt.Println(rover.message())
}