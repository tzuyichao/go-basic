package main

import (
    "fmt"
    "reflect"
)

func main() {
    one := [2]string {"bar", "foo"}
    var two []string

    fmt.Println(reflect.TypeOf(one))
    fmt.Println(reflect.TypeOf(two))
}