package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] == val {
			nums[i] = 1000
		} else {
			l += 1
		}
	}
	sort.Ints(nums)
	return l
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	len := removeElement(nums, 2)
	fmt.Println("len:", len)
	fmt.Println(nums)
}