package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(44))
}

func climbStairs(n int) int {
	var memo []int = make([]int, n)
	return climb(0, n, memo)
}

func climb(i int, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = climb(i+1, n, memo) + climb(i+2, n, memo)
	return memo[i]
}
