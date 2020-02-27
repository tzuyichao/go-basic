package main

import (
    "fmt"
    "sync"
)

type visited struct {
	mu sync.Mutex
	visited map[string]int
}

func (v *visited) VisitLink(url string)int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
    
}