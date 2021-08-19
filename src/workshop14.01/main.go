package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	r, err := http.Get("https://www.google.com")
	if err != nil {
	    log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}