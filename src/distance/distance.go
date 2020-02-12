package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var distance = rand.Intn(401000000-56000000) + 56000000
    fmt.Printf("distance of Mars and Earch: %v", distance)
}