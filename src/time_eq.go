package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

func main() {
    t1 := time.Now()
    data, err := json.Marshal(t1)
    if err != nil {
        log.Fatal(err)
    }
    var t2 time.Time
    if err := json.Unmarshal(data, &t2); err != nil {
        log.Fatal(err)
    }
    fmt.Println(t1 == t2)
}