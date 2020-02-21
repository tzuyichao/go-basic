package main

import (
    "fmt"
)

type Planets []string

func terraform(planets Planets) {
    for i := range planets {
        planets[i] = "New " + planets[i]
	}
	fmt.Printf("terraform: %v\n", planets)
}

func main() {
    our := [...]string {
        "Mercury", "Venus", "Earth", "Mars",
        "Jupiter", "Saturn", "Uranus", "Neptune",
    }
    var solar_system Planets = Planets(our[:])
    fmt.Println(solar_system)
    terraform(solar_system)
    fmt.Println(solar_system)
}