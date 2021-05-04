package main

import (
    "fmt"
)

func main() {
    myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<- myChannel)
}

func greeting(ch chan string) {
	ch <- "hi"
}