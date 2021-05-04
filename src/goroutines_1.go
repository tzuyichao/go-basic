package main

import (
    "fmt"
	"time"
)

func main() {
    go a()
    go b()
	time.Sleep(time.Second)
    fmt.Println("end main()")
}

func a() {
    for i := 0; i < 50; i++ {
        fmt.Print("a")
    }
}

func b() {
    for i := 0; i < 50; i++ {
        fmt.Print("b")
    }
}