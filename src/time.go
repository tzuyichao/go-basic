package main

import (
	"fmt"
	"time"
)

func main() {
	// type inference timeout type is int
	timeout := 3
	// var timeout time.Duration = 3
	// const timeout = 3
	fmt.Printf("before ")
	time.Sleep(timeout * time.Millisecond)
	// time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("after")
}