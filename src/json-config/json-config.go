package main

import (
    "fmt"
    "os"
    "encoding/json"
)

type configuration struct {
	Enabled bool
	Path string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
	fmt.Println(conf.Enabled)
}