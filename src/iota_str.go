package main

import (
    "fmt"
)

type FilePerm uint16

const (
    Read FilePerm = 1 << iota
    Write
    Execute
)

func (p FilePerm) String() string {
    switch p {
    case Read:
        return "read"
    case Write:
        return "write"
    case Execute:
        return "execute"
    }
    return fmt.Sprintf("Unknown FilePerm: %d", p)
}

func main() {
    fmt.Println(Execute)
    fmt.Printf("%d\n", Execute)
}