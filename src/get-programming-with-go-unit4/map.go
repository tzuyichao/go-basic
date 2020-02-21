package main

import "fmt"

func main() {
    temperature := map[string]int {
		"Earth": 15,
		"Mars": -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the earth is %v°C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	// key does not exist, get 0 for int type
	temp = temperature["Moon"]
	fmt.Printf("On average the earth is %v°C.\n", temp)
	if temp, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the earth is %v°C.\n", temp)
	} else {
		fmt.Println("Moon not found")
	}
}