package main

import (
    "time"
    "fmt"
)

func whatstheclock() string {
    return time.Now().Format(time.ANSIC)
}

func main() {
    fmt.Println(whatstheclock())
}