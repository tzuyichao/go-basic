package main

import (
	"fmt"
)

func main() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true, "Ted": false}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}