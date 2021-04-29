package main

import (
    "fmt"
)

func main() {
    array1 := [5]string{"a", "b", "c", "d", "e"}
    slice1 := array1[0:3]

    fmt.Println(array1)
    fmt.Println(slice1)

    array1[1] = "X"
    fmt.Println(array1)
    fmt.Println(slice1)

    slice1[1] = "x"
    fmt.Println(array1)
    fmt.Println(slice1)
}