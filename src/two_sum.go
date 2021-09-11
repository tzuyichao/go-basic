package main

import (
	"fmt"
	"sort"
)

func solve(list []int, target int) (int, int, error) {
	sort.Ints(list)
	low, high := 0, len(list)-1

	for low < high {
		sum := list[low] + list[high]
		if sum < target {
			low++
		} else if sum > target {
			high--
		} else if sum == target {
			return list[low], list[high], nil
		}
	}

	return 0, 0, fmt.Errorf("target %d not found", target)
}

func main() {
	data := []int{5, 1, 3, 6}
	x, y, err := solve(data, 9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x, y)
}