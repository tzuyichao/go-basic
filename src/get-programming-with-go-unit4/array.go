package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(planets)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	// invalid array index 8 (out of bounds for 8-element array)
	// compile error因為compiler可以檢查的到
	// 如果寫法compiler檢查不到則runtime會引發panic
	// planets[8] = "Pluto"

	dwarfs := [5]string {"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i:=0; i<len(dwarfs); i++ {
		fmt.Println(dwarfs[i])
	}

	status := [...]string {"draft", "open", "deleted"}
	fmt.Println(len(status))
	for i, elem := range status {
		fmt.Printf("%d: %v\n", i, elem)
	}

	status2 := status
	status2[1] = "Woops"
	fmt.Println(status)
	fmt.Println(status2)
}