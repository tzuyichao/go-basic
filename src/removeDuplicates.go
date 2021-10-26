package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j<len(nums); j++ {
		if nums[i] != nums[j] {
			i = i + 1
			nums[i] = nums[j]
		}
	}
	return i+1
}

/**
 Leetcode 26 easy
 */
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len := removeDuplicates(nums)
	fmt.Println(len, nums)
}