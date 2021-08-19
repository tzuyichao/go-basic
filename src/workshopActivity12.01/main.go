package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type txnCategory string

const (
	fuel txnCategory = "fuel"
	food txnCategory = "food"
	mortgage txnCategory = "mortgage"
	repairs txnCategory = "repairs"
	insurance txnCategory = "insurance"
	utilities txnCategory = "utilities"
	retirement txnCategory = "retirement"
)

type transaction struct {
	id string
	payee string
	spent float32
	category txnCategory
}

func main() {
	file, err := os.Open("bank.csv")
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
		fmt.Println(record[3])
		categoryItem := strings.Trim(record[3], " ")
		switch(categoryItem) {
		case string(fuel):
			fmt.Println("Got Fuel")
		default:
			fmt.Println("Else")
		}
	}
}