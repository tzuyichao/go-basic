package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Num CPU:", runtime.NumCPU())
}