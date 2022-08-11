package main

import (
    "fmt"
)

func main() {
    m := map[string]int{
        "hello": 3,
    }

    key := []byte{'h', 'e', 'l', 'l', 'o'}
    val := m[string(key)]       // no memory allocation
    fmt.Println(val)
}