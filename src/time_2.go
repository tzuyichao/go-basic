package main

import (
    "fmt"
    "time"
)

func main() {
    day := time.Now().Weekday()
    hour := time.Now().Hour()
    fmt.Println("Day:", day, "Hour:", hour)
    if day.String() == "Monday" {
        if hour >= 1 {
            fmt.Println("Performing full blown test!")
        } else {
            fmt.Println("Performing hit-n-run test!")
        }
    } else {
        fmt.Println("Performing hit-n-run test!")
    }
}