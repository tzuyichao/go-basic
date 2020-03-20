package main

import (
    "runtime"
)

type data struct {
    i, j int
}

// GODEBUG=gctrace=1 go run sliceGC.go
func main() {
	var N = 40000000
	var structure []data
	for i := 0; i<N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	_ = structure[0]
}