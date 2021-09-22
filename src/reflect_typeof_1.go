package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}

func main() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
}