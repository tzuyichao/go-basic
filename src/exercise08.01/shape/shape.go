package shape

import "fmt"

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base float64
	Height float64
}

type Rectangle struct {
	Length float64
	Width float64
}

type Square struct {
	Side float64
}

func PrintShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("%s的面積: %.2f\n", item.name(), item.area())
	}
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) name() string {
	return "三角形"
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "長方形"
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) name() string {
	return "正方形"
}