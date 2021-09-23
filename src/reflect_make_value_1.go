package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	var stringSliceType = reflect.TypeOf([]string(nil))
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	fmt.Println("cap: ",  ssv.Cap(), ", len:", ssv.Len())
	ss := ssv.Interface().([]string)
	fmt.Println("data:", ss, ", cap: ", cap(ss), ", len:", len(ss))

	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")

	ssv = reflect.Append(ssv, sv)
	ss = ssv.Interface().([]string)
	fmt.Println("data:", ss, ", cap: ", cap(ss), ", len:", len(ss))
}