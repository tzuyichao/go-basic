package main

import (
    "fmt"
    "strconv"
)

func main() {
    countdown := 10
    str := "Lauch in T minus " + strconv.Itoa(countdown) + " seconds."
    fmt.Println(str)
}