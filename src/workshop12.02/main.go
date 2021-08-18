package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "your first name")
	i := flag.Int("age", -1, "your age")
	b := flag.Bool("married", false, "are you married?")
	flag.Parse()

	if *n == "" {
		fmt.Println("Name is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Name:    ", *n)
	fmt.Println("Age:     ", *i)
	fmt.Println("Married: ", *b)
}