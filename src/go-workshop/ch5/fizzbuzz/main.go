package main

import (
    "fmt"
)

func main() {
	fmt.Println("main is in control")
	fizzBuzz()
	fmt.Println("Back to main")
}

func fizzBuzz() {
	for i := 0; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}