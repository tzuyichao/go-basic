package main

import (
	"fmt"
)

func main() {
	const a = `⌘`

	fmt.Printf("plain string: %s\n", a)
	fmt.Printf("quoted string: %q\n", a)

	fmt.Printf("hex bytes: ")
	for i:=0; i<len(a); i++ {
		fmt.Printf("%x", a[i])
	}
	fmt.Printf("\n")
}