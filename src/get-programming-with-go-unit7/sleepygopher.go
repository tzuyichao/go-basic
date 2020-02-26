package main

import (
    "fmt"
    "time"
)

func main() {
	go sleepyGopher()
	fmt.Println("main go sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Done.")
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("....snore....")
}