package main

import (
    "fmt"
)

func clamDown() {
    fmt.Println(recover())
}

func freakOut() {
    defer clamDown()
    panic("oh no")
}

func main() {
    freakOut()
    fmt.Println("Exiting Normally")
}