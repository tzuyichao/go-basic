package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Println("Hello, World!")

	unix.Syscall(unix.SYS_WRITE, 1, uintptr(unsafe.Pointer(&[]byte("Hello, World!")[0])), uintptr(len("Hello, World!")))
}
