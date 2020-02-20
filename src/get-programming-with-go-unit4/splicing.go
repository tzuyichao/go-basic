package main

import "fmt"

func main() {
    planets := [...]string {
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	iceGiants2 := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants)
	fmt.Println(iceGiants2)

	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)
}