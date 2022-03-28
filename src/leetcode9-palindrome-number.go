package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		s := strconv.Itoa(x)
		r := len(s) - 1
		l := 0
		for l < r {
			if s[l] != s[r] {
				return false
			}
			r -= 1
			l += 1
		}
		return true
	}
}

/**
 * https://leetcode.com/problems/palindrome-number/
 *
 * Runtime: 12 ms, faster than 91.88% of Go online submissions for Palindrome Number.
 * Memory Usage: 4.6 MB, less than 60.93% of Go online submissions for Palindrome Number.
 */
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(0))
}
