package main

import (
    "fmt"
    "encoding/json"
    "os"
)

type greeting struct {
    SomeMessage string `json:"message"`
}

func main() {
    data := []byte(`
    {
        "message": "Greeting fellow gopher!
    }
    `)
    if !json.Valid(data) {
        fmt.Printf("JSON is not valid: %s\n", data)
        os.Exit(1)
    } else {
        fmt.Println("JSON is valid")
    }
}