package main

import (
    "fmt"
    "sort"
)

func main() {
    planets := []string {
		"Mercury", "Venus", "Earth", "Mars",
        "Jupiter", "Saturn", "Uranus", "Neptune",
	}
	fmt.Println(planets)
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	saturn_idx := sort.StringSlice(planets).Search("Saturn")
	fmt.Println(saturn_idx)
	// 根據文件寫的
	// The return value is the index to insert x if x is not
    // present (it could be len(a)).
	test_idx := sort.StringSlice(planets).Search("Test")
	fmt.Println(test_idx)
}