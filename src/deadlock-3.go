package main

import (
  "fmt"
)

func main() {
  c := make(chan string)
  c <- "message"
  fmt.Println(<- c)
}
