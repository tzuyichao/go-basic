package main

import (
	"fmt"
	"log"
	"sync"
)

func work(from, to int, wg *sync.WaitGroup, res *string, mtx *sync.Mutex) {
	for i:=from; i<=to; i++ {
		mtx.Lock()
		*res = fmt.Sprintf("%s|%d|", *res, i)
		mtx.Unlock()
	}
	wg.Done()
	return
}

func main() {
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	res := ""
	go work(1, 25, wg, &res, mtx)
	go work(26, 50, wg, &res, mtx)
	go work(51, 75, wg, &res, mtx)
	go work(76, 100, wg, &res, mtx)
	wg.Wait()
	log.Println(res)
}