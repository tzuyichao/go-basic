package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"time"
)

type Order struct {
	ID string `json:id`
	DateOrdered time.Time `json:date_ordered`
	CustomerId string `json:customer_id`
	Items []Item `json:items`
}

type Item struct {
	ID string `json:id`
	Name string `json:name`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
	}
	var order1 = Order{
		ID: "001",
		DateOrdered: time.Now(),
		CustomerId: "C002",
		Items: []Item{
			Item{
				ID: "i001",
				Name: "The Go Programming Language",
			},
			Item{
				ID: "i002",
				Name: "Learning Go",
			},
		},
	}
	data, err := json.Marshal(order1)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ioutil.WriteFile(os.Args[1], data, 0644); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Done.")
}