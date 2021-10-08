package main

import (
	"time"
	"fmt"
    "os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: convertTimes parse_string")
		return
	}
	dateString := os.Args[1]

	now, err := time.Parse("02 January 2006 15:04 MST", dateString)
	if err == nil {
		fmt.Println("now:", now)
		loc, _ := time.LoadLocation("America/New_York")
		fmt.Printf("New York Time: %s\n", now.In(loc))
    } else {
		fmt.Println("err:", err)
	}
}