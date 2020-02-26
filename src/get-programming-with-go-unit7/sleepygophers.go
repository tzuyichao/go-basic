package main

import (
    "fmt"
    "time"
)

func main() {
	for i:=0; i<5; i++ {
	    go sleepyGopher(i)
    }
	fmt.Println("main go sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Done.")
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("....%d snore....\n", id)
}