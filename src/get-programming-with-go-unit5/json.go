package main

import (
    "fmt"
    "encoding/json"
    "os"
)

func main() {
	// Fields name should begin with upercase
    type location struct {
		Lat, Long float64
	}

	curiosity := location{-4.5895, 137.4417}
	
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}