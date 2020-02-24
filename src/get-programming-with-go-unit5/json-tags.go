package main

import (
    "fmt"
    "encoding/json"
    "os"
)

func main() {
    type location struct {
		Lat float64 `json:"latitude"`
		Long float64 `json:"longitude"`
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