package main

import (
    "fmt"
    "runtime"
)

func main() {
	fmt.Printf("You are using %v on %v machine\n", runtime.Compiler, runtime.GOARCH)
	fmt.Printf("Using Go version %v\n", runtime.Version())
	fmt.Printf("Number of CPUs: %v\n", runtime.NumCPU()) 
    fmt.Printf("Number of Goroutines: %v\n", runtime.NumGoroutine()) 
}