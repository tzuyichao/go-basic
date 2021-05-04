package main

import (
    "fmt"
)

func greeting(ch chan string) {
    // ch <- "hi"
	fmt.Println("in goroutine")
	ch <- "hi"
}

func main() {
    myChannel := make(chan string)
	// myChannel <- "hi from main"
    go greeting(myChannel)
    //greeting(myChannel)
    fmt.Println(<- myChannel)
}