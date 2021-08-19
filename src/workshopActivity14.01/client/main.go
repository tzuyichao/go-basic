package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Names struct {
	Names []string `json:"names"`
}

func getDataAndReturnResponse() (electricCount int, boogalooCount int) {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := Names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}
	electricCount = 0
	boogalooCount = 0
	for _, name := range names.Names {
		switch(name) {
		case "Electric":
			electricCount += 1
		case "Boogaloo":
			boogalooCount += 1
		}
	}
	return electricCount, boogalooCount
}

func main() {
	eCount, bCount := getDataAndReturnResponse()
	log.Printf("Electric: %d\n", eCount)
	log.Printf("Boogaloo: %d\n", bCount)
}