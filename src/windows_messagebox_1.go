package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")
	messagebox = moduser32.NewProc("MessageBoxW")
)

const (
	NULL = 0
	MB_OK = 0
)

func MessageBox(title, caption string) int {
	ret, _, _ := messagebox.Call(
		uintptr(NULL), 
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))), 
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
	    uintptr(MB_OK))
	return int(ret)
}

func main() {
	ret := MessageBox("Hello", "It works!")
	fmt.Println("ret:", ret)
}