package main

import (
	"context"
	"log"
	"time"
)

func countNumbers(c context.Context, out chan int) {
	v := 0
	for {
		select {
		case <-c.Done():
			out <- v
			break
		default:
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}

func main() {
	r := make(chan int)
	c := context.TODO()
	c1, stop := context.WithCancel(c)
	go countNumbers(c1, r)

	go func() {
		time.Sleep(time.Millisecond * 100 * 3)
		stop()
	}()
	v := <- r
	log.Println(v)
}