package main

import (
    "fmt"
    "os"
	"log"
)

func main() {
    wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("starting dir:", wd)

	if err := os.Chdir("/"); err != nil {
		log.Fatal(err)
		return
	}

	if wd, err = os.Getwd(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("final dir:", wd)
}