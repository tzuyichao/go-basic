package main

import (
	"fmt"
)

func SelectionSort(src []int) []int {
	var res []int = make([]int, len(src), cap(src))

	copy(res, src)

	for i := 0; i < len(res)-1; i++ {
		minIdx := i
		for j := minIdx+1; j < len(res); j++ {
			if res[minIdx] > res[j] {
				minIdx = j
			}
		}
		if i != minIdx {
			res[i], res[minIdx] = res[minIdx], res[i]
		}
	}

	return res
}

func main() {
	data := []int{4, 3, 2, 1}
	res := SelectionSort(data)
	fmt.Printf("%v\n", res)
}