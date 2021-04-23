package main

import (
    "fmt"
    "reflect"
)

func main() {
    var test_int int = 5
	var test_int32 int32 = 5
	var test_int64 int64 = 5
    fmt.Println("test_int:", reflect.TypeOf(test_int))
    fmt.Println("test_int32:", reflect.TypeOf(test_int32))
    fmt.Println("test_int64:", reflect.TypeOf(test_int64))
}