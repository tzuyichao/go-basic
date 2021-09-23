package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int `myTag:"value"`
	B string `myTag:"value2"`
	C string `myTag:"10"`
}

func main() {
	var f Foo
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		currentField := ft.Field(i)
		fmt.Println(currentField.Name, currentField.Type.Name(), currentField.Tag.Get("myTag"))
	}
}

