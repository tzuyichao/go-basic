package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	firstName = iota
	lastName
	age
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := true
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
		if header {
			header = false
			continue
		}
		fmt.Println("-----------------")
		fmt.Println("First Name:", record[firstName])
		fmt.Println("Last Name :", record[lastName])
		fmt.Println("Age       :", record[age])
	}
}