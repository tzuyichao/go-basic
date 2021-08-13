package main

import (
	"log"
	"time"
)

func sum(from, to int) int {
	res := 0
	for i:=from; i<=to; i++ {
		res += i
	}
	return res
}

func main() {
	log.Println("Go")
	var s1, s2, s3 int
	go func() {
		s1 = sum(1, 100)
	}()
	time.Sleep(time.Second)
	s2 = sum(1, 10)
	go func() {
		s3 = sum(1, 10)
	}()
	log.Println(s1, s2, s3)
}