package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	result := make([]string, len(worlds))
	for i := range worlds {
		result[i] = prefix + " " + worlds[i]
	}
	return result
}

func main() {
	dwarfs := terraform("new", "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Printf("length %v capacity %v %v", len(dwarfs), cap(dwarfs), dwarfs)
}