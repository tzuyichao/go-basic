package main

import (
	"fmt"
	"sort"
)

type result struct {
	val1 int
	val2 int
}

func solve(list []int, target int) []result {
	sort.Ints(list)
	low, high := 0, len(list)-1
	r := []result{}
	for low < high {
		left := list[low]
		right := list[high]
		sum := left + right
		if sum < target {
			low++
		} else if sum > target {
			high--
		} else {
			r = append(r, result{left, right})
			for low < high && list[low] == left {
				low++
			}
			for low < high && list[high] == right {
				high--
			}
		}
	}
	return r
}

func main() {
	nums := []int{1, 3, 1, 2, 2, 3}
	results := solve(nums, 4)
	for _, r := range results {
		fmt.Println(r)
	}
}