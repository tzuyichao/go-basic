package main

import (
	"os"
	"log"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	log.Println(f)

	f.Write([]byte("使用Write()寫入\n"))
	f.WriteString("使用WriteString()寫入\n")
}