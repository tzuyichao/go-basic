package main

import (
    "log"
    "io/ioutil"
    "os"
)

func main() {
	filename  := os.Args[1]
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Println(err)
        return
    }

    text := string(contents)
    log.Println(text)

}