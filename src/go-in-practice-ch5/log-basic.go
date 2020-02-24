package main

import (
    "log"
)

func main() {
    log.Println("This is a regular message.")
    log.Fatalln("This is a fatal error.")
    log.Println("You cannoe see this log message due to previous Fatalln() called")
}