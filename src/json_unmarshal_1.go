package main

import (
    "fmt"
    "encoding/json"
)

type greeting struct {
    SomeMessage string `json:"message"`
}

func main() {
	data := []byte(`
	{
		"message": "Greetings fellow gopher!"
	}
	`)
	var g greeting
	err := json.Unmarshal(data, &g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.SomeMessage)
}