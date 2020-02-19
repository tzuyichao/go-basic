package main

import (
    "fmt"
)

type getRow func(row int) (string, string)

func drawTable(row int, r getRow) {
    x, y := r(row)
    fmt.Printf("%v, x=%v, y=%v\n", row, x, y)
}

func getFakeRow(row int) (string, string) {
    return "1", "2"
}

func main() {
    var r getRow = getFakeRow
    drawTable(10, r)
}