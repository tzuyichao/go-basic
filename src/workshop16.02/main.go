package main

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i:=from; i<=to; i++ {
		*res+=i
	}
	wg.Done()
	return
}

func main() {
	s1 := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 100, &wg, &s1)
	wg.Wait()
	log.Println(s1)

	s2 := 0
	s3 := 0
	wg2 := sync.WaitGroup{}
	wg2.Add(2)
	go sum(1, 10, &wg2, &s2)
	go sum(100, 150, &wg2, &s3)
	wg2.Wait()
	log.Println(s2, s3)
}