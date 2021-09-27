package main

import (
	"fmt"
	"unsafe"
)

func state() {
	var onload = createEvents("onload")
	var receive = createEvents("receive")
	var success = createEvents("success")

	mapEvents := make(map[string]interface{})
	mapEvents["messageOnload"] = unsafe.Pointer(onload)
	mapEvents["messageReceive"] = unsafe.Pointer(receive)
	mapEvents["messageSuccess"] = unsafe.Pointer(success)

	fmt.Println(*(*string)(mapEvents["messageReceive"].(unsafe.Pointer)))
	fmt.Println(*(*string)(mapEvents["messageSuccess"].(unsafe.Pointer)))
}

func createEvents(s string) *string {
	return &s
}

func main() {
	state()
}