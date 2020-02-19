package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Printf("%6.2f° K is %6.2f° C\n", k, c)
	c2 := k.celsius()
	fmt.Printf("%6.2f° K is %6.2f° C\n", k, c2)
}