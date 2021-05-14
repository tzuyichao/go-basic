package main

import (
    "greeting"
)

// https://blog.golang.org/go116-module-changes#TOC_3.
func main() {
    greeting.Hello()
    greeting.Hi()
}