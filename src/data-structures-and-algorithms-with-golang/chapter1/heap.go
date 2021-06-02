package main

import (
	"fmt"
	"container/heap"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap

	n = len(previous)
	x1 = previous[n-1]

	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1,4,5}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}