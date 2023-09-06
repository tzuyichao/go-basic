package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := makeRandomSlice(10, 100)
	printSlice(slice, 10)
	checkSorted(slice)
	bubbleSort(slice)
	checkSorted(slice)
}
