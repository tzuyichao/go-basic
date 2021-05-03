package main

import (
    "fmt"
)

func freakOut() {
    panic("oh no")
    recover()
}

func main() {
    freakOut()
    fmt.Println("Exiting normally")
}