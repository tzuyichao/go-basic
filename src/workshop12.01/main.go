package main

import (
	"flag"
	"fmt"
)

func main() {
	v := flag.Int("value", -1, "Need a value for the flag.")
	flag.Parse()
	fmt.Println(v)
	fmt.Println(*v)
}

