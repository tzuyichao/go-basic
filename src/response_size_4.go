package main

import (
    "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
    URL string
    Size int
}

func responseSize(url string, ch chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- Page{URL: url, Size:len(body)}
}

func main() {
    urls := [3]string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	ch := make(chan Page)
	for _, url := range urls {
		go responseSize(url, ch)
	}
	for i:=0; i<len(urls); i++ {
		page := <- ch
		fmt.Printf("%s:%d\n", page.URL, page.Size)
	}
}