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

func print(out chan int, wg *sync.WaitGroup) {
	for i := range out {
		log.Println(i)
	}
	wg.Done()
	return
}

func main() {
	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	out := make(chan int)
	wg.Add(2)
	wg2.Add(1)
	go source("./input1.dat", out, wg)
	go source("./input2.dat", out, wg)
	go print(out, wg2)
	wg.Wait()
	close(out)
	wg2.Wait()
}