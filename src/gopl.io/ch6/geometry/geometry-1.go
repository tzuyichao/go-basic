package main

import (
	"math"
	"fmt"
)

type Point struct {X, Y float64}

func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) dummy() {
	p.X = 0.0
	p.Y = 0.0
}

func (p *Point) zero() {
	p.X = 0.0
	p.Y = 0.0
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim)
	fmt.Println(perim.Distance())

	p := Point {1, 1}
	fmt.Println(p)
	p.dummy()
	fmt.Println(p)
	p.zero()
	fmt.Println(p)
}