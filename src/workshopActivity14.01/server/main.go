package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type server struct {}
type Names struct {
	Names []string `json:names`
}

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	names := Names{}

	for i := 0; i<rand.Intn(5)+1; i++ {
		names.Names = append(names.Names, "Electric")
	}
	for i := 0; i<rand.Intn(5)+1; i++ {
		names.Names = append(names.Names, "Boogaloo")
	}
	jsonBytes, _ := json.Marshal(names)
	log.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}