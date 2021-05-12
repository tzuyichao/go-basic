package main

import (
    "crypto/rand"
    "fmt"
)

func main() {
    buf := make([]byte, 10)
    _, err := rand.Read(buf)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(fmt.Sprintf("%x", buf))
}