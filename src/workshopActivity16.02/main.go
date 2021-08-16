package main

import (
	"log"
    "sync"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func source(file string, out chan int, wg *sync.WaitGroup) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				wg.Done()
				return
			} else {
				panic(err)
			}
		}
		trimStr := strings.ReplaceAll(str, "\r\n", "")
		i, err := strconv.Atoi(trimStr)
		if err != nil {
			panic(err)
		}
		out <- i
	}
}

func split(in, odd, even chan int, wg *sync.WaitGroup) {
	for i := range in {
		log.Println(i)
		switch i%2 {
		case 0:
			even <- i
		case 1:
			odd <- i
		}
	}
	close(even)
	close(odd)
	wg.Done()
	return
}

func sum(in, out chan int, wg *sync.WaitGroup) {
    sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
	wg.Done()
}

func merge(odd, even chan int, wg *sync.WaitGroup) {
	oddsum := <- odd
	evensum := <- even
	log.Printf("Odd: %d\n", oddsum)
	log.Printf("Even: %d\n", evensum)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	out := make(chan int)
	odd := make(chan int)
	even := make(chan int)
	sumodd := make(chan int)
	sumeven := make(chan int)
	wg.Add(2)
	wg2.Add(4)
	go source("./input1.dat", out, wg)
	go source("./input2.dat", out, wg)
	go split(out, odd, even, wg2)
	go sum(odd, sumodd, wg2)
	go sum(even, sumeven, wg2)
	go merge(sumodd, sumeven, wg2)
	wg.Wait()
	close(out)
	wg2.Wait()
}