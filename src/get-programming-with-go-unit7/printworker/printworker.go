package main

import (
	"fmt"
	"time"
)

func worker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <- next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func main() {
	go worker()
	time.Sleep(4 * time.Second)
}