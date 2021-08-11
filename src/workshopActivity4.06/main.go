package main

import (
	"fmt"
)

func checker(v interface{}) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float32, float64:
		return "float"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Printf("1 is %v\n", checker(1))
	fmt.Printf("3.14 is %v\n", checker(3.14))
	fmt.Printf("hello is %v\n", checker("hello"))
	fmt.Printf("true is %v\n", checker(true))
	fmt.Printf("{} is %v\n", checker(struct{}{}))
}