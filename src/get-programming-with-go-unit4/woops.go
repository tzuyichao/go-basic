package main

import "fmt"

func main() {
	planets := map[string]string {
		"Earth", "Sector ZZ9",
		"Mars", "Sector ZZ9",
	}

	planetMarkII = planets
	planets["Earth"] = "Woops"
	fmt.Println(planets)
	fmt.Println(planetMarkII)

	delete(planets, "Earth")
	fmt.Println(planetMarkII)
}