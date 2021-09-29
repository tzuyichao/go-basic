package main

import (
	"fmt"
	"context"
	"time"
	"sync"
)

type key struct{}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i:=0; i<5; i++ {
		go func(ctx context.Context) {
			v := ctx.Value(key{})
			fmt.Println("key:", v)
			wg.Done()
			fmt.Println("Done ", v)
			<-ctx.Done()
			fmt.Println("Exit: ", ctx.Err(), v)
		}(context.WithValue(ctx, key{}, i))
	}
	wg.Wait()
	cancel()
	time.Sleep(time.Second)
}