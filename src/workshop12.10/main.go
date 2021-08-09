package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(record)
	}
}