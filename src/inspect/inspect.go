package main

import (
	"fmt"
	"time"
)

func main() {
	year := 2008
	var magic int64 = 2010200920082007
	days := 365.2424
	text := "test"
	test := true
	var red, green, blue uint8 = 0, 141, 213

	fmt.Printf("Type: %T for %v\n", year, year)
	fmt.Printf("Type: %T for %v\n", magic, magic)
	fmt.Printf("Type: %T for %[1]v\n", days, days)
	fmt.Printf("Type: %T for %v\n", text, text)
	fmt.Printf("Type: %T for %v\n", test, test)
	fmt.Printf("Type: %T for %v\n", red, red)
	fmt.Printf("%x %x %x\n", red, green, blue)
	fmt.Printf("background: #%02x%02x%02x\n", red, green, blue)
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	blue = 255
	fmt.Printf("blue: %08b\n", blue)
	blue++
	fmt.Printf("blue: %08b\n", blue)

	future := time.Unix(12622780800, 0)
	fmt.Printf("Type: %T for %v\n", future, future)
	fmt.Println(future)
}