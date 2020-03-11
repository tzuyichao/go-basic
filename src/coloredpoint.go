package main

import (
    "image/color"
    "fmt"
)

type Point struct {
    X, Y float64
}

type ColoredPoint struct {
    Point
    Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Printf("%v\n", cp)
}