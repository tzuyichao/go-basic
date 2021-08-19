package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:message`
}

func postDataAndReturnResponse(msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)
	r, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
    if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}

func main() {
	msg := messageData{Message: "Hi Server"}
	data := postDataAndReturnResponse(msg)
	fmt.Println(data)
}