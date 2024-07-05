package main

import (
  "fmt"
  "sync"
)

func main() {
  fmt.Println("Total Items Packed:", PackItems(0))
}

func PackItems(totalItems int) int {
  const workers = 2
  const itemsPerWorker = 1000

  var wg sync.WaitGroup
  itemsPacked := 0
  for i := 0; i < workers; i++ {
    wg.Add(1)
    go func(workerID int) {
      defer wg.Done()
      for j := 0; j < itemsPerWorker; j++ {
        itemsPacked = totalItems
	itemsPacked++
	totalItems = itemsPacked
      }
    } (i)
  }

  wg.Wait()
  return totalItems
}
