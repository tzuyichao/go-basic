package main

import (
	"fmt"
	"sort"
)

func findK(nums []int) int {
	k := 0
	for k = 0; k < len(nums); k++ {
		if k+1 < len(nums) {
			if nums[k+1] < nums[k] {
				return k+1
			}
		}
	}
	return k
}

func makeOrigin(nums []int, k int) []int {
	res := make([]int, len(nums))

	for i:=0; i<len(nums); i++ {
		res[i] = nums[(i+k)%len(nums)]
	}

	return res
}

func search(nums []int, target int) int {
	k := findK(nums)
	ori := makeOrigin(nums, k)
	l := len(nums)
	idx := sort.SearchInts(ori, target)
	if ori[idx] != target {
		return -1
	}
	return (idx+k)%l
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	r := search(nums, 0)
	fmt.Println(r)

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	r = search(nums, 3)
	fmt.Println(r)

	nums = []int{1}
	r = search(nums, 0)
	fmt.Println(r)

	nums = []int{4, 5, 6, 7, 8, 0}
	r = search(nums, 6)
	fmt.Println(r)
}