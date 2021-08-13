package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func sumWithAtomic(from, to int, wg *sync.WaitGroup, res *int32) {
	for i:=from; i<=to; i++ {
		atomic.AddInt32(res, int32(i))
	}
	wg.Done()
	return
}

func sumWithoutAtomic(from, to int, wg *sync.WaitGroup, res *int32) {
	for i:=from; i<=to; i++ {
		*res = *res + int32(i)
	}
	wg.Done()
	return
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go sumWithAtomic(1, 25, wg, &s1)
	go sumWithAtomic(26, 50, wg, &s1)
	go sumWithAtomic(51, 75, wg, &s1)
	go sumWithAtomic(76, 100, wg, &s1)
	wg.Wait()
	log.Println(s1)

	// s2 := int32(0)
	// wg2 := &sync.WaitGroup{}
	// wg2.Add(4)
	// go sumWithoutAtomic(1, 25, wg2, &s2)
	// go sumWithoutAtomic(26, 50, wg2, &s2)
	// go sumWithoutAtomic(51, 75, wg2, &s2)
	// go sumWithoutAtomic(76, 100, wg2, &s2)
	// wg2.Wait()
	// log.Println(s2)
}