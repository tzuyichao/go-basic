package main

import (
	"fmt"
	"context"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/json"
)

type NEResult struct {
	Company []string `json:"company"`
	Person []string `json:"person"`
	Problem []string `json:"problem"`
	Property []string `json:"property"`
	Area []string `json:"area"`
	Product []string `json:"product"`
	Machine []string `json:"machine"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("api url is required.")
		return
	}
	apiUrl := os.Args[1]
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	data := "<img-245>如上圖，請說明如何在人機介面操作螢幕上做一個透明按鍵，使隱藏的ON/OFF/EXIT鍵，按下透明按鍵時顯示供操作；操作完成後按EXIT鍵即恢復隱藏。"
	req, err := http.NewRequestWithContext(
		context.Background(), 
		http.MethodPost, 
		apiUrl, 
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "text/plain, application/json, application/*+json, */*")
	fmt.Println(req)
	fmt.Println(req.Body)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Println(res)
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", resData)
	var ne = NEResult{}
	//err = json.NewDecoder(res.Body).Decode(&ne)
	err = json.NewDecoder(bytes.NewReader(resData)).Decode(&ne)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ne)
}