package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%6.2f° K is %6.2f° C\n", kelvin, celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%6.2f° C is %6.2f° F\n", celsius, fahrenheit)
}