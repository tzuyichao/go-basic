package main

import "fmt"

func dump(label string, slice []string) {
    fmt.Printf("%v: length %v capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string {"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// [start index, length and data index, capacity index]
	dwarfs1 := dwarfs[2:3:4]
	dump("dwarfs1", dwarfs1)
}