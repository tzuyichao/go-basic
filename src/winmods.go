package main

import (
	"syscall"
	"fmt"
)

var (
	ModKernel32 = syscall.NewLazyDLL("kernel32.dll")
)

func main() {

}