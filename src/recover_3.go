package main

import (
    "fmt"
)

func clamDown() {
    recover()
}

func freakOut() {
    defer clamDown()
    panic("oh no")
}

func main() {
    freakOut()
    fmt.Println("Exiting Normally")
}