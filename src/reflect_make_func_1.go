package main

import (
	"fmt"
	"time"
	"reflect"
)

// from Learning Go
func MakeTimedFunction(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperFunc := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		end := time.Now()
		fmt.Println("Elapsed time:", end.Sub(start))
		return out
	})
	return wrapperFunc.Interface()
}

func main() {
    f1 := MakeTimedFunction(func (target int) int {
		var r int = 0
		for i:=0; i<target; i++ {
			r += i
		}
		time.Sleep(1)
		return r
	}).(func(int)int)
	fmt.Println(f1(100000))
}