package main

import "fmt"

func main() {
	const distance = 236e15
	const lightSpeed = 299792
	const secondsPerYear = 31536000

	fmt.Println("Canis Major Dwarf's light year to Earch is", distance/lightSpeed/secondsPerYear)
}