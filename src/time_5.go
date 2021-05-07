package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    only_after, _ := time.Parse(time.RFC3339, "2032-11-01T22:08:41+00:00")
    fmt.Println(now, only_after)
    fmt.Println(now.After(only_after))
    if now.After(only_after) {
        fmt.Println("Executing actions!")
    } else {
        fmt.Println("Now is not the time yet!!!")
    }
}