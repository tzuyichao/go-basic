package main

import (
    "fmt"
)

type celsius float64

type fahrenheit float64

type header func()

type footer func()

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c*9.0/5.0)+32.0)
}

type body func()

func drawTable(h header, b body, f footer) {
	h()
	b()
	f()
}

func drawHeader() {
	fmt.Println("=======================")
	fmt.Println("| °C       | °F       |")
	fmt.Println("=======================")
}

func drawFooter() {
	fmt.Println("=======================")
}

func makeDrawBody(start celsius, end celsius, step celsius) body {
	return func() {
		for t:=start; t<=end; t+=step {
			fmt.Printf("| %5.1f    | %5.1f    |\n", t, t.fahrenheit())
		}
	}
}

func main() {
	var drawBody body = makeDrawBody(celsius(-40), celsius(100), celsius(5))
	drawTable(drawHeader, drawBody, drawFooter)
}