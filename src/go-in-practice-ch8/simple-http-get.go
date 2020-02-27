package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	res, err := http.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Printf("%v", err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	res.Body.Close()
	fmt.Printf("%s", b)
}