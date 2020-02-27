package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <- c:
			fmt.Printf("gopher %d has finished sleeping\n", gopherId)
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Printf("... %d snore ...\n", id)
	ch <- id
}