package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    responseSize("https://example.com/")
    responseSize("https://golang.org/")
    responseSize("https://golang.org/doc")
}

func responseSize(url string) {
    fmt.Println("Getting", url)
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}