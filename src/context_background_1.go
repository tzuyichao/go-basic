package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	done := ctx.Done()
	for i:=0; ; i++ {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}

