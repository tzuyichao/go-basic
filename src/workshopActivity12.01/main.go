package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"errors"
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

func mapTxnCategory(categoryItem string) (txnCategory, error) {
	switch(categoryItem) {
	case string(fuel):
		return fuel, nil
	case string(food):
		return food, nil
	case string(mortgage):
		return mortgage, nil
	case string(repairs):
		return repairs, nil
	case string(insurance), "car insurance", "life insurance":
		return insurance, nil
	case string(utilities):
		return utilities, nil
	case string(retirement):
		return retirement, nil
	default:
		return "", errors.New(fmt.Sprintf("given %s category not found", categoryItem))
	}
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
		
		categoryItem := strings.Trim(record[3], " ")
		category, err := mapTxnCategory(categoryItem)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s => %v\n", categoryItem, category)
	}
}