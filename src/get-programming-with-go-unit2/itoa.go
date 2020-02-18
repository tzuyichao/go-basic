package main

import (
    "fmt"
    "strconv"
)

func main() {
    countdown := 10
    str := "Lauch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)
	
	str2 := fmt.Sprintf("Lauch in T minus %v seconds.", countdown)
	fmt.Println(str2)
	
	size, err := strconv.Atoi("10a")
	if err != nil {
		fmt.Println("ascii to integer failed")
	}
	fmt.Println(size)
}