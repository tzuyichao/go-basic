package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string {"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
	dump("dwarfs[2:4]", dwarfs[2:4])
	// slice[n1:n2]後的slice的capacity的計算是到原本slice的長度
	dwarfs1 := append(dwarfs, "Orcus")
	dump("dwarfs", dwarfs)
	dump("dwarfs1", dwarfs1)
}