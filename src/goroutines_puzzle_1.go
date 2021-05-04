package main

import (
    "fmt"
    "time"
)

func main() {
    go repeat("x")
    go repeat("y")
    time.Sleep(time.Second)
}

func repeat(s string) {
    for i:=0; i<25; i++ {
        fmt.Print(s)
    }
}