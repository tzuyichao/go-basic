package main

import "fmt"

func main() {
    dwarfs := []string {"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
    dwarfs = append(dwarfs, "Orcus", "Quaoar")
	fmt.Println(dwarfs)
	
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
	solr := planets[:]
	solr = append(solr, "Kitty")
	fmt.Println(planets)
	fmt.Println(solr)	
}