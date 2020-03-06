package main

import (
    "fmt"
)

func main() {
	// anonymous struct and struct literal
    data := struct {
		Name string
	} {
		Name: "text1",
	}

	fmt.Println(data)

	datas := []struct {
		Name string
	} {
		{
			Name: "title1",
		},
		{
			Name: "title2",
		},
		{
			Name: "title3",
		},
	}

	fmt.Println(datas)
}