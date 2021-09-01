package main

import (
	"fmt"
)

func InsertionSort(src []int) []int {
	var res = make([]int, len(src))

	copy(res, src)

	for i:=1; i<len(res); i++ {
		key := res[i]
		j := i - 1
		for j>=0 && res[j]>key {
			res[j+1] = res[j]
			j--
		}
		res[j+1] = key
	}

	return res
}

func main() {
	data := []int{4, 3, 2, 1}
	res := InsertionSort(data)
	fmt.Println(res)
}