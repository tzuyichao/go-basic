package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "os"
)

func postFileAndReturnResponse(filename string) string {
	fileDataBuffer := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&fileDataBuffer)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	formFile, err := multipartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	multipartWriter.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080", &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	fmt.Println(postFileAndReturnResponse("main.go"))
}