package main

import "fmt"

func main() {
    fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
	
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`
	peace be upon you
	upon you be peace`)
}