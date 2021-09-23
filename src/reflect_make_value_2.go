package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int
	it := reflect.TypeOf(i)
	fmt.Println(it.Name())
	iv := reflect.New(it).Elem()
	iv.SetInt(10)
	
	fmt.Println(iv)
}