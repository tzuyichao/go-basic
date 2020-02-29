package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n", c1)
}