package main

import (
    "fmt"
    "time"
)

func reportNap(name string, delay int) {
    for i:=0; i<delay; i++ {
        fmt.Println(name, "sleeping")
        time.Sleep(1 * time.Second)
    }
    fmt.Println(name, "wakes up!")
}

func send(ch chan string) {
    reportNap("sending gorouting", 2)
    fmt.Println("*** sending value ***")
    ch <- "a"
    fmt.Println("*** sending value ***")
    ch <- "b"
}

func main() {
    myChannel := make(chan string)
    go send(myChannel)
    reportNap("receiving gorouting", 5)
    fmt.Println(<- myChannel)
    fmt.Println(<- myChannel)
}