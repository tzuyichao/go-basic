package main

import (
    "fmt"
)

func main() {
    array1 := [5]string{"a", "b", "c", "d", "e"}
    slice1 := array1[0:3]

    fmt.Println(array1)
    fmt.Println(slice1)

    slice2 := append(slice1, "c")
    fmt.Println("array1: ", array1)
    fmt.Println("slice1: ", slice1)
    fmt.Println("slice2: ", slice2)

	// 故意超過
	slice2 = append(slice2, "c")
	slice2 = append(slice2, "c")
	fmt.Println("array1: ", array1)
    fmt.Println("slice2: ", slice2)

	// 此時跟slice2和array1已無關係
	slice2[0] = "X"
	fmt.Println("array1: ", array1)
    fmt.Println("slice2: ", slice2)

	array1[1] = "X"
	fmt.Println("array1: ", array1)
	fmt.Println("slice1: ", slice1)
    fmt.Println("slice2: ", slice2)
}