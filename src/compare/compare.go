package main

import "fmt"

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	
	var age = 44
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	var apple = "apple"
	var banana = "banana"
	var greater = apple > banana

	fmt.Printf("is 'apple' greater than 'banana'? %v",greater)
}