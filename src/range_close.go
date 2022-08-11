package main

import (
    "fmt"
)

func fibs(n int) chan int {
    ch := make(chan int)
    go func() {
		defer close(ch)
        a, b := 1, 1
        for i:=0; i<n; i++ {
            ch <- a
            a, b = b, a+b
        }
    }()
    return ch
}

func main() {
    for i := range fibs(5) {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}