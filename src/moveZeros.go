package main

import (
  "fmt"
  "sort"
)

func moveZeroes(nums []int)  {
    sort.SliceStable(nums, func(i, j int) bool {
        if nums[i] == 0 {
            return false
        }
        if nums[j] == 0 {
            return true
        }
        return i < j
    })
}

func main() {
  arr := []int{3, 0, 0, 1, 2, 0, 0, 0, 4}

  moveZeroes(arr)

  fmt.Println(arr)
}
