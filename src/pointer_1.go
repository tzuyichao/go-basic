package main

import (
    "fmt"
    "reflect"
)

func main() {
	fmt.Println("Type of Pointer")
    var myInt int
	fmt.Println(reflect.TypeOf(&myInt))

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
	
	fmt.Println("Pointer Data")
	myInt = 4
	myIntPointer := &myInt
	fmt.Println("&myInt:", &myInt)
	fmt.Println("myIntPointer:", myIntPointer)
	fmt.Println("*myIntPointer:", *myIntPointer)
	myInt = 8
	fmt.Println("myIntPointer:", myIntPointer)
	fmt.Println("*myIntPointer:", *myIntPointer)
}