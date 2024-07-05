package main

import (
  "log"
  "sync"
  "runtime"
)

func main() {
  times := 0
  for {
    times++
    counter := PackItems(0)
    if counter != 2000 {
      log.Fatalf("it should be 2000 but found %d on execution %d", counter, times)
    }
  }
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
	runtime.Gosched()
	itemsPacked++
	totalItems = itemsPacked
      }
    } (i)
  }

  wg.Wait()
  return totalItems
}
