package main

/**
 * 2. Add Two Numbers
 * https://leetcode.com/problems/add-two-numbers/
 *
 * Runtime: 12 ms, faster than 67.48% of Go online submissions for Add Two Numbers.
 * Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Add Two Numbers.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func len(l *ListNode) int {
	r := 0
	c := l
	for c != nil {
		c = c.Next
		r += 1
	}

	return r
}

func carry(l *ListNode) *ListNode {
	c := l
	isCarry := false
	c_val := 0
	for c != nil {
		if isCarry {
			c.Val += c_val
		}
		if c.Val >= 10 {
			isCarry = true
			c_val = c.Val / 10
		} else {
			isCarry = false
			c_val = 0
		}
		c.Val = c.Val % 10
		if c.Next == nil && isCarry {
			c.Next = &ListNode{0, nil}
		}
		c = c.Next
	}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := l1
	r2 := l2

	c1 := r1
	c2 := r2
	len1 := len(c1)
	len2 := len(c2)
	if len1 >= len2 {
		for c1 != nil {
			if c2 != nil {
				c1.Val += c2.Val
				c2 = c2.Next
			}
			c1 = c1.Next
		}
		return carry(r1)
	} else {
		for c2 != nil {
			if c1 != nil {
				c2.Val += c1.Val
				c1 = c1.Next
			}
			c2 = c2.Next
		}
		return carry(r2)
	}
}

func print(l *ListNode) {
	c := l
	for c != nil {
		println("> ", c.Val)
		c = c.Next
	}
	println()
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	print(addTwoNumbers(l1, l2))
	println("Done.")
}
