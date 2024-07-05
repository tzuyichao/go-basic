package main

import (
  "fmt"
  "time"
  "sync"
)

func main() {
  wg := sync.WaitGroup{}
  wg.Add(2)
  go say("hello", &wg)
  go say("world", &wg)
  wg.Wait()
}

func say(s string, wg *sync.WaitGroup) {
  defer wg.Done()
  for i := 1; i < 5; i++ {
    time.Sleep(500 * time.Millisecond)
    fmt.Println(s)
  }
}
