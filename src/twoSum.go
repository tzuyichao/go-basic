package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	store := make(map[int]int)

	for idx, val := range nums {
		found, ok := store[val]
		if ok {
			res[0] = found
			res[1] = idx
			return res
		} else {
			store[target-val] = idx
		}
	}

	return res
}

func main() {
	var res []int
	res = twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}