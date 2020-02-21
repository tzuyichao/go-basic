// https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html
package main

import (
	"fmt"
	"unsafe"
)

func InspectSlice(slice []string) {
	address := unsafe.Pointer(&slice)

	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	addPtr := (*[10]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr [0x%x]\n", address, lenAddr, capAddr)
	fmt.Printf("Slice Length[%d] Cap[%d]\n", *lenPtr, *capPtr)

	fmt.Printf("First 10 Elemnts\n")
	for index := 0; index < *lenPtr; index++ {
		if index < *lenPtr {
		    fmt.Printf("[%d] %p %s\n", index, &(*addPtr)[index], (*addPtr)[index])
	    }
	}
}

func main() {
	//dwarfs := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	dwarfs := make([]string, 8)
	InspectSlice(dwarfs)
}