package main

import (
	"time"
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	time.AfterFunc(time.Second*10, cancel)
    done := ctx.Done()
	for i:=0; ; i++ {
		select {
		case <-done:
			fmt.Println("Exit: ", ctx.Err())
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}