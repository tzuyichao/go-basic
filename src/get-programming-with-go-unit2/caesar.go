package main

import "fmt"

func main() {
    c := 'a'
    c = c + 3
	fmt.Printf("%c %[1]v\n", c)
	
	c = 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c\n", c)

	c = 'A'
	fmt.Printf("A: %v\n", c)

	for i := 'A'; i <= 'z'; i++ {
		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println()

	str := "L fdph, L vdz, L frqtxhuhg"
	for _, ch := range str {
		if ch != ' ' && ch != ',' {
		    fmt.Printf("%c", (ch-3))
	    } else {
			fmt.Printf("%c", ch)
		}
	}
}