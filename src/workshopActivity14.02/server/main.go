package main 

import (
	"log"
	"net/http"
	"encoding/json"
)

var names []string

type server struct{}

type messageData struct {
	Names []string `json:names`
}

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "POST":
		jsonDecoder := json.NewDecoder(r.Body)
		messageData := messageData{}
		err := jsonDecoder.Decode(&messageData)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(messageData.Names)
		for _, name := range messageData.Names {
			names = append(names, name)
		}
		log.Println(names)
		w.Write([]byte("POST"))
	case "GET":
		messageData := messageData{}
		for _, name := range names {
			messageData.Names = append(messageData.Names, name)
		}
		jsonBytes, _ := json.Marshal(messageData)
		w.Write(jsonBytes)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}