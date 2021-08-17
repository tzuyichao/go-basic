package main

import (
	"fmt"
	"sync"
)

type Workers struct {
	in, out chan int
	subWorkers int
	mtx sync.Mutex
}

func (w *Workers) init(maxSubWorkers, maxData int) {
	w.in, w.out = make(chan int, maxData), make(chan int)
	w.mtx = sync.Mutex{}
	for i:=0; i<maxSubWorkers; i++ {
		w.mtx.Lock()
		w.subWorkers++
		w.mtx.Unlock()
		go w.readThem()
	}
}

func (w *Workers) readThem() {
	sum := 0
	for i := range w.in {
		sum += i
	}
	w.out <- sum
	w.mtx.Lock()
	w.subWorkers--
	w.mtx.Unlock()
	if w.subWorkers <= 0 {
		close(w.out)
	}
}

func (w *Workers) addData(item int) {
	w.in <- item
}

func (w *Workers) getherResult() int {
	close(w.in)
	total := 0
	for i := range w.out {
		total += i
	}
	return total
}

func main() {
    maxWorkers := 10
	maxData := 100
	workers := Workers{}
	workers.init(maxWorkers, maxData)
	for i:=1; i<=maxData; i++ {
		workers.addData(i)
	}
	res := workers.getherResult()
	fmt.Println(res)
}