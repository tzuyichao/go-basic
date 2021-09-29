package main

import (
	"context"
	"time"
	"fmt"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	time.AfterFunc(time.Second*10, cancel)
	done := ctx.Done()
	for i:=0; ; i++ {
		select {
		case <-done:
			fmt.Println("exit:", ctx.Err())
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}