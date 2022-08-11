package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 1)
    ch <- 1
    a, ok := <-ch
    fmt.Println(a, ok)
    close(ch)
    b, ok := <-ch
    fmt.Println(b, ok)
}