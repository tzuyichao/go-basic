package main

import "fmt"

func main() {
	var hours = 28 * 24
	var distance = 56000000
	var speed = distance / hours
	fmt.Println("speed", speed, "km/h")
}