package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

type server struct {}

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadFile, uploadFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadFile.Close()
	fileContent, err := ioutil.ReadAll(uploadFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(uploadFileHeader)
	log.Println(string(fileContent))
	w.Write([]byte("Done"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}