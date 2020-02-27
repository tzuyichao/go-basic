package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i:=0; i<5; i++ {
		go sleepyGopher(i, c)
	}
	for i:=0; i<5; i++ {
		gopherId := <- c
		fmt.Printf("gopher %d has finished sleeping\n", gopherId)
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Printf(".... %d snore ....\n", id)
	c <- id
}