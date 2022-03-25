package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	// var test int = 9646324351
	//  fmt.Println(test)
	var res int
	var isNeg = x < 0
	x = int(math.Abs(float64(x)))
	idx := 1
	l := int(math.Floor(math.Log10(float64(x))))
	for x/10 > 0 {
		m := x % 10
		//fmt.Println(m, res)
		res += m * int(math.Pow10(l))
		l -= 1
		idx += 1
		x = x / 10
	}
	// fmt.Println(res)
	res += x
	// fmt.Println(res)
	if math.Log2(float64(res)) > 31 {
		return 0
	}
	if isNeg {
		res = -res
	}
	// fmt.Println(res)
	return res
}

func main() {
	fmt.Println(reverse(123))
}
